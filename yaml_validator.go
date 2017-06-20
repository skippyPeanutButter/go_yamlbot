package main

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

}
