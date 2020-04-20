package runtime

import (
	"../core"
	"../parser"
	"log"
)

var functions = map[string]func(...interface{}){
	"SMÅGLI":      core.Println,
	"FULLSPÄCKAD": core.Printf,
	"TILLGÅNG":    core.Save,
}

var getters = map[string]func(...interface{}) string {
	"SMÅKALLT": core.Get,
}

func Run(ast []parser.Node) []parser.Node {
	for i := 0; i < len(ast); i++ {
		node := ast[i]

		// Execute the found method
		if node.Type == "CallExpression" {
			// Move the params in an array of string
			var params []interface{}
			for _, param := range node.Params {
				if param.Type == "CallExpression" {
					for _, p := range Run([]parser.Node{param}) {
						params = append(params, p.Value)
					}

					continue
				}

				params = append(params, param.Value)
			}

			// Search an existent function in the map
			function := functions[node.Value]
			getter := getters[node.Value]
			if function == nil && getter == nil {
				log.Fatal("This function was not found.")
			}

			if function != nil {
				// Execute the found function
				function(params...)
			}

			if getter != nil {
				ast[i] = parser.Node{
					Type: "String",
					Value: getter(params...),
				}
			}
		}

		i++
	}

	return ast
}
