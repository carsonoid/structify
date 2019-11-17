package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"

	"github.com/ghodss/yaml"
	"github.com/sanity-io/litter"
	typepkg "k8s.io/api/{{ .APIVersion }}"
)

func main() {
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	var o typepkg.{{ .Kind }}
	if err := yaml.Unmarshal(b, &o); err != nil {
		panic(err)
	}

	// extra litter config
	sq := litter.Options{
		StripPackageNames: false,
	}
	pr := sq.Parse(o)

	result, err := template.New("result").ParseFiles("templates/result.go.tpl")
	if err != nil {
		panic(err)
	}

	var r bytes.Buffer
	err = result.ExecuteTemplate(&r, "result.go.tpl", pr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", r.String())
}
