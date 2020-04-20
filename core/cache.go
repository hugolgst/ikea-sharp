package core

import "fmt"

var cache = map[string]string{}

func Save(params ...interface{}) {
	cache[fmt.Sprintf("%v", params[0])] = fmt.Sprintf("%v", params[1])
}

func Get(params ...interface{}) string {
	return cache[fmt.Sprintf("%v", params[0])]
}
