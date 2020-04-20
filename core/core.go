package core

var functions = map[string]func(...interface{}){}
var getters = map[string]func(...interface{}) string{}

// GetFunctions returns the map of functions
func GetFunctions() map[string]func(...interface{}) {
	return functions
}

// GetGetters returns the map of getters
func GetGetters() map[string]func(...interface{}) string {
	return getters
}

// AddFunction links the given name with the given function in the functions map
func AddFunction(name string, fn func(...interface{})) {
	functions[name] = fn
}

// AddGetter links the given name with the given getter in the getters map
func AddGetter(name string, fn func(...interface{}) string) {
	getters[name] = fn
}
