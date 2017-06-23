// Utility methods
package main

func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

func merge(map1 map[interface{}]interface{}, map2 map[interface{}]interface{}) map[interface{}]interface{} {
	for k, v := range map2 {
		map1[k] = v
	}
	return map1
}

func toInterfaceArray(strarr []string) []interface{} {
	new := make([]interface{}, len(strarr))
	for i, v := range strarr {
		new[i] = v
	}
	return new
}
