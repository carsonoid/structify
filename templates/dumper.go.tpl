package main

import (
	"io/ioutil"
	"os"

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

	litter.Dump(o)
}
