package runtime

import (
	"../core"
	"../log"
	"../parser"
)

// Run takes a given AST and runs the found functions and getters
func Run(ast []parser.Node) []parser.Node {
	for i := 0; i < len(ast); i++ {
		node := ast[i]

		// Execute the found method
		if node.Type == "CallExpression" {
			params := GetParams(node)

			// Search an existent function in the map
			function, getter := core.GetFunctions()[node.Value], core.GetGetters()[node.Value]
			if function == nil && getter == nil {
				log.Errorf("This reference was not found.")
			}

			// Execute the found function
			if function != nil {
				function(params...)
				// Execute the found getter and replace the content in the AST
			} else if getter != nil {
				ast[i] = parser.Node{
					Type:  "String",
					Value: getter(params...),
				}
			}
		}

		i++
	}

	return ast
}

// GetParams retrieves the params from a given node and returns the params as an array of interface
func GetParams(node parser.Node) (params []interface{}) {
	for _, param := range node.Params {
		// Execute recursion if we found another CallExpression
		if param.Type == parser.CallExpression {
			for _, p := range Run([]parser.Node{param}) {
				params = append(params, p.Value)
			}

			continue
		}

		params = append(params, param.Value)
	}

	return
}
