package main

import (
	"fmt"
	"io/ioutil"

	"github.com/ghodss/yaml"
	"github.com/sanity-io/litter"
	typepkg "k8s.io/api/{{ .APIVersion }}"
)

func main() {
	y, err := ioutil.ReadFile("{{ .FilePath }}")
	if err != nil {
		panic(err)
	}

	var o typepkg.{{ .Kind }}
	err = yaml.Unmarshal(y, &o)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	litter.Dump(o)
}
