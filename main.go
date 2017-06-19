package main

import (
	"github.com/smallfish/simpleyaml"
	"io/ioutil"
)

func main() {
	source, err := ioutil.ReadFile(".yamlbot.yml")
	if err != nil {
		panic(err)
	}
	y, err := simpleyaml.NewYaml(source)
	if err != nil {
		panic(err)
	}

	validateRules(y)
}
