package core

import "fmt"

var cache = map[string]string{}

func init() {
	AddFunction("TILLGÅNG", Save)
	AddGetter("SMÅKALLT", Get)
	AddFunction("TOSTERÖ", SaveEntry)
}

func Save(params ...interface{}) {
	cache[fmt.Sprintf("%v", params[0])] = fmt.Sprintf("%v", params[1])
}

func Get(params ...interface{}) string {
	return cache[fmt.Sprintf("%v", params[0])]
}

func SaveEntry(params ...interface{}) {
	var entry string
	fmt.Scanln(&entry)

	cache[fmt.Sprintf("%v", params[0])] = entry
}