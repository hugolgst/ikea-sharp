package traverser

import "../parser"

type Visitor map[string]func(node *parser.Node, parent parser.Node)

// The traverser will execute the functions found in the AST
func Traverser(ast parser.Node, visitor Visitor) {
	TraverseNode(ast, parser.Node{}, visitor)
}

// TraverseArray will iterate through the given nodes and execute the next function
func TraverseArray(body []parser.Node, parent parser.Node, visitor Visitor) {
	for _, child := range body {
		TraverseNode(child, parent, visitor)
	}
}

// TraverseNode will search for a function to execute and execute it, and execute TraverseArray if
// the current node cannot be executed.
func TraverseNode(node, parent parser.Node, visitor Visitor) {
	// Iterate through the visitor elements and if it matches the node one, executes the function
	for functionType, function := range visitor {
		if functionType == node.Type {
			function(&node, parent)
		}
	}

	switch node.Type {
	case "Program":
		TraverseArray(node.Body, node, visitor)
		break

	case "CallExpression":
		TraverseArray(node.Params, node, visitor)
		break

	case "NumberLiteral":
		break
	}

}
