package main

import (
	"bufio"
	"os"

	"github.com/ghodss/yaml"
	"github.com/sanity-io/litter"
	typepkg "k8s.io/api/{{ .APIVersion }}"
)

func main() {
	s := bufio.NewScanner(os.Stdin)

	s.Scan()

	var o typepkg.{{ .Kind }}
	if err := yaml.Unmarshal(s.Bytes(), &o); err != nil {
		panic(err)
	}

	litter.Dump(o)
}
