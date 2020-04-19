package parser

import "../tokenizer"

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
func Parse(tokens []tokenizer.Token) Node {
	var currentIndex int

	// Create the default program node
	program := Node{
		Type: "PROGRAM",
		Body: []Node{},
	}

	// Iterate through the tokens
	for currentIndex < len(tokens) {

	}

	return program
}