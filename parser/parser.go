package parser

import (
	"../tokenizer"
)

var (
	tokens       []tokenizer.Token
	currentIndex int
)

// A Node is an element in the AST
type Node struct {
	Type       string
	Value      string
	Name       string
	Callee     *Node
	Expression *Node
	Body       []Node
	Params     []Node
	Arguments  *[]Node
	Context    *[]Node
}

// Parse takes the given tokens and creates an AST
func Parse(_tokens []tokenizer.Token) Node {
	tokens = _tokens

	// Create the default program node
	program := Node{
		Type: "PROGRAM",
		Body: []Node{},
	}

	// Iterate through the tokens
	for currentIndex < len(tokens) {
		program.Body = append(program.Body, Iterate())
	}

	return program
}

// Iterate will browse the tokens and returns a node
func Iterate() Node {
	token := tokens[currentIndex]

	// Return a number literal node if the current token is a number
	if token.Type == "NUMBER" {
		currentIndex++

		return Node{
			Type:  "NumberLiteral",
			Value: token.Value,
		}
	}

	// Iterate through the parameters inside the parentheses
	if token.Type == "PARENTHESES" && token.Value == tokenizer.LeftParentheses {
		// Skip the parenthese and get to the parameters
		currentIndex++
		token := tokens[currentIndex]

		// Begin to crete the CallExpression node by adding the name
		node := Node{
			Type:   "CallExpression",
			Name:   token.Value,
			Params: []Node{},
		}

		// Skip the name
		currentIndex++
		token = tokens[currentIndex]

		// Add parameters until we get the right parentheses
		for token.Type != "PARENTHESES" && token.Value != tokenizer.RightParentheses {
			node.Params = append(node.Params, Iterate())
			token = tokens[currentIndex]
		}

		// Skip the last parentheses and returns the node
		currentIndex++
		return node
	}
	
	return Node{}
}
