package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"text/template"

	"github.com/ghodss/yaml"
)

func main() {
	filePath := os.Args[1]

	y, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	var obj map[string]interface{}
	var objInterface interface{}
	err = yaml.Unmarshal(y, &obj)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(y, &objInterface)
	if err != nil {
		panic(err)
	}

	var apiVersion string
	if v, ok := obj["apiVersion"]; ok {
		apiVersion, ok = v.(string)
		if !ok {
			panic(fmt.Errorf("Cannot convert apiVersion to string"))
		}
	} else {
		panic(fmt.Errorf("No apiVersion in Object"))
	}

	var kind string
	if v, ok := obj["kind"]; ok {
		kind, ok = v.(string)
		if !ok {
			panic(fmt.Errorf("Cannot convert kind to string"))
		}
	} else {
		panic(fmt.Errorf("No kind in Object"))
	}

	fmt.Printf("Dumping %s/%s\n", apiVersion, kind)

	// Special cases for apiVersion
	if apiVersion == "v1" {
		apiVersion = "core/v1"
	}

	dynamicDump(apiVersion, kind, y)
}

type dumperInput struct {
	APIVersion string
	Kind       string
}

func dynamicDump(apiVersion string, kind string, data []byte) {

	tmpl, err := template.New("dumper").ParseFiles("templates/dumper.go.tpl")
	if err != nil {
		panic(err)
	}

	var b bytes.Buffer

	err = tmpl.ExecuteTemplate(&b, "dumper.go.tpl", dumperInput{
		APIVersion: apiVersion,
		Kind:       kind,
	})
	if err != nil {
		panic(err)
	}

	ioutil.WriteFile("dumpers/test.go", b.Bytes(), 0600)

	// Run generated dumper. Pass data via stdin
	fmt.Println("Running generated dumper")
	dumper := exec.Command("go", "run", "dumpers/test.go")

	stdin, err := dumper.StdinPipe()
	if err != nil {
		panic(err)
	}
	defer stdin.Close()

	dumper.Stdout = os.Stdout
	dumper.Stderr = os.Stderr

	if err := dumper.Start(); err != nil {
		panic(err)
	}

	stdin.Write(data)

	dumper.Wait()
}
