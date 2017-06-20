package main

import (
	"errors"
	"fmt"
	"github.com/smallfish/simpleyaml"
	"reflect"
	"strings"
)

func validateRules(y *simpleyaml.Yaml) {
	if y.Get("defaults").IsFound() {
		defaults, err := y.Get("defaults").Map()
		if err != nil {
			fmt.Println("The 'defaults' section must be a yaml map type")
			panic(err)
		}
		validateRulesKeys(defaults)
	}

	if !y.Get("rules").IsFound() {
		err := errors.New("Rules section not defined in .yamlbot.yml")
		panic(err)
	}
	validateEachKey(y)
}

func validateEachKey(y *simpleyaml.Yaml) {
	var defaults map[interface{}]interface{}
	if y.Get("defaults").IsFound() {
		defaults, _ = y.Get("defaults").Map()
	}
	if !y.Get("rules").IsArray() {
		err := "The rules section of a rules file must define a list of keys."
		panic(err)
	}
	arr, _ := y.Get("rules").Array()
	for _, keymap := range arr {
		if len(defaults) == 0 {
			validateKey(keymap.(map[interface{}]interface{}))
		} else {
			merged_map := merge(defaults, keymap.(map[interface{}]interface{}))
			validateKey(merged_map)
		}
	}
}

func validateKey(keymap map[interface{}]interface{}) {
	validateRulesKeys(keymap)
	validateKeyExistsAndIsString(keymap)
	validateRequiredKeyExistsAndIsBool(keymap)
	fmt.Println("Key: " + keymap["key"].(string) + " validated.")
}

func validateRulesKeys(keymap map[interface{}]interface{}) {
	validkeys := []string{
		"key",
		"required_key",
		"and_requires",
		"or_requires",
		"value_whitelist",
		"value_blacklist",
		"types",
	}
	var invalidkeys []string
	for k, _ := range keymap {
		if !contains(validkeys, k.(string)) {
			invalidkeys = append(invalidkeys, k.(string))
		}
	}

	if len(invalidkeys) == 0 {
		return
	} else {
		msg := "Invalid key(s) specified in rules file: " + strings.Join(invalidkeys, " ") + "\n"
		msg += "Valid rules keys include: " + strings.Join(validkeys, " ") + "\n"
		panic(msg)
	}
}

func validateKeyExistsAndIsString(keymap map[interface{}]interface{}) {
	if keymap["key"] == nil || reflect.ValueOf(keymap["key"]).Kind() != reflect.String {
		msg := "Missing required key 'key' within rules file.\n"
		msg += "Or a key name has a value that is not a String.\n"
		panic(msg)
	}
}

func validateRequiredKeyExistsAndIsBool(keymap map[interface{}]interface{}) {
	if keymap["required_key"] == nil || reflect.ValueOf(keymap["required_key"]).Kind() != reflect.Bool {
		msg := "Missing required key 'required_key' for key: " + keymap["key"].(string) + "\n"
		msg += "Or 'required_key' has a value that is not a Boolean.\n"
		panic(msg)
	}
}
