package main

import (
	"fmt"
	"strings"
)

var violations uint

func validateYaml(rules *simpleyaml.Yaml, yamlfile *simpleyaml.Yaml) uint {
	violations = 0
	rules, _ := y.Get("rules").Array()
	var defaults map[interface{}]interface{}

	if y.Get("defaults").IsFound() {
		defaults, _ := y.Get("defaults").Map()
	}

	for _, rule := range rules {
		violations += validateBasedOnRule(rule.(map[interface{}]interface{}), yamlfile, defaults)
	}
	return violations
}

func validateBasedOnRule(rule map[interface{}]interface{}, y *simpleyaml.Yaml, defaults map[interface{}]interface{}) uint {
	invalid := false
	rule = merge(defaults, rule)
	y.GetPath(toInterfaceArray(rule))
	checkIfKeyIsRequired(rule, y)
	checkAndRequires()
	checkOrRequires()
	checkValWhiteList()
	checkValBlackList()
	checkTypes()
	if invalid {
		return 1
	} else {
		return 0
	}
}

func checkIfKeyIsRequired(rule map[interface{}]interface{}, y *simpleyaml.Yaml) bool {
	key, err := rule.Get("key").String()
	if err != nil {
		panic(err)
	}
	value, err := y.GetPath(toInterfaceArray(strings.Split(key, "."))...).String()
	if err != nil {
		panic(err)
	}
	if (!rule.Get("required_key").IsFound() || rule.Get("required_key").Bool()[0] == false) && value == nil {
		return false
	}
	if !rule.Get("or_requires").IsFound() {
		return false
	}

	fmt.Println("VIOLATION: Missing required key: " + key)
}

func checkAndRequires() {

}
func checkOrRequires() {

}

func checkValWhiteList() {

}

func checkValBlackList() {

}

func checkTypes() {

}
