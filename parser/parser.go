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
	Type   string
	Value  string
	Params []Node
}

// Parse takes the given tokens and creates an AST
func Parse(_tokens []tokenizer.Token) []Node {
	tokens = _tokens

	// Create the default program node array
	var program []Node

	// Iterate through the tokens
	for currentIndex < len(tokens) {
		node := Iterate()

		if node.Type != "" {
			program = append(program, node)
		}
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
	if token.Value == tokenizer.LeftParentheses {
		call := Node{
			Type:   "CallExpression",
			Value:  tokens[currentIndex-1].Value,
			Params: []Node{},
		}

		for tokens[currentIndex+1].Value != tokenizer.RightParentheses {
			currentIndex++
			token = tokens[currentIndex]

			if tokens[currentIndex+1].Value == tokenizer.LeftParentheses {
				currentIndex++

				call.Params = append(call.Params, Iterate())
				return call
			}

			call.Params = append(call.Params, Node{
				Type: token.Type,
				Value: token.Value,
			})
		}
		currentIndex++

		return call
	}

	if token.Type == "SEMICOLON" {
		currentIndex++
		return Node{
			Type: "SemiColon",
			Value: token.Value,
		}
	}

	currentIndex++

	return Node{}
}
