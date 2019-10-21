package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"text/template"

	"github.com/ghodss/yaml"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	convert     = kingpin.Command("convert", "Convert one or more Kubernetes documents to structs")
	convertPath = convert.Arg("FILE", "Path to document(s) to convert").String()

	serve     = kingpin.Command("serve", "Start Structify service")
	serveAddr = serve.Flag("addr", "Listen address for server").Default(":8080").String()
)

func main() {
	switch kingpin.Parse() {
	// Command line conversion
	case "convert":
		convertFile(*convertPath)

	// HTTP Server
	case "serve":
		startServer()
	}
}

func startServer() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/structify", structifyHandler)

	fmt.Printf("Starting listener on %s\n", *serveAddr)
	http.ListenAndServe(*serveAddr, nil)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}

// StructifyRequest takes serialized data for a k8s object
// and any config options
// TODO: Config Options
type StructifyRequest struct {
	Data string
}

// StructifyResponse hold the resulting struct and metadata
type StructifyResponse struct {
	Struct string `json:"struct"`
}

func convertFile(filePath string) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	ss := getStructBytes(data)

	fmt.Fprint(os.Stderr, "--------------------------\n")
	fmt.Printf("%s", ss)
}

func structifyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusMethodNotAllowed)
		return
	}

	if r.Body == nil {
		http.Error(w, "No data", 400)
		return
	}

	var sr StructifyRequest
	err := json.NewDecoder(r.Body).Decode(&sr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	output, err := json.Marshal(StructifyResponse{
		Struct: string(getStructBytes([]byte(sr.Data))),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}

func getStructBytes(data []byte) []byte {
	apiVersion, kind, err := getAPIVersionAndKind(data)
	if err != nil {
		panic(err)
	}

	writeDumper(apiVersion, kind)

	results, err := getDumperResults(data)
	if err != nil {
		panic(err)
	}

	return results
}

func getAPIVersionAndKind(data []byte) (string, string, error) {
	var obj map[string]interface{}
	if err := yaml.Unmarshal(data, &obj); err != nil {
		return "", "", err
	}

	var apiVersion string
	if v, ok := obj["apiVersion"]; ok {
		apiVersion, ok = v.(string)
		if !ok {
			return "", "", fmt.Errorf("Cannot convert apiVersion to string")
		}
	} else {
		return "", "", fmt.Errorf("No apiVersion in Object")
	}

	var kind string
	if v, ok := obj["kind"]; ok {
		kind, ok = v.(string)
		if !ok {
			return "", "", fmt.Errorf("Cannot convert kind to string")
		}
	} else {
		return "", "", fmt.Errorf("No kind in Object")
	}

	fmt.Fprintf(os.Stderr, "Dumping %s/%s\n", apiVersion, kind)

	// Special cases for apiVersion
	if apiVersion == "v1" {
		apiVersion = "core/v1"
	}

	return apiVersion, kind, nil
}

type dumperInput struct {
	APIVersion string
	Kind       string
}

func getDumperResults(data []byte) ([]byte, error) {
	fmt.Fprintf(os.Stderr, "Running generated dumper\n")
	dumper := exec.Command("go", "run", "dumpers/test.go")

	// stdin comes from given data
	dumper.Stdin = bytes.NewBuffer(data)

	// save stdout and stderr to a buffer
	var stdout, stderr bytes.Buffer
	dumper.Stdout = &stdout
	dumper.Stderr = &stderr

	err := dumper.Run()
	if err != nil {
		return stdout.Bytes(), err
	}

	if stderr.Len() != 0 {
		return stdout.Bytes(), fmt.Errorf("%s", stderr.Bytes())
	}

	return stdout.Bytes(), nil
}

func writeDumper(apiVersion, kind string) error {
	tmpl, err := template.New("dumper").ParseFiles("templates/dumper.go.tpl")
	if err != nil {
		return err
	}

	var b bytes.Buffer
	err = tmpl.ExecuteTemplate(&b, "dumper.go.tpl", dumperInput{
		APIVersion: apiVersion,
		Kind:       kind,
	})
	if err != nil {
		return err
	}
	ioutil.WriteFile("dumpers/test.go", b.Bytes(), 0600)

	return nil
}
