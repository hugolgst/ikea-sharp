package parser

import (
	"../tokenizer"
	"strings"
)

const (
	CallExpression = "CallExpression"
	NumberLiteral  = "NumberLiteral"
	SemiColon      = "SemiColon"
	String         = "String"
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
			Type:  NumberLiteral,
			Value: token.Value,
		}
	}

	// Iterate through the parameters inside the parentheses
	if token.Value == tokenizer.LeftParentheses {
		call := Node{
			Type:   CallExpression,
			Value:  tokens[currentIndex-1].Value,
			Params: []Node{},
		}

		// Iterate until we found the closing parentheses
		for tokens[currentIndex+1].Value != tokenizer.RightParentheses {
			currentIndex++
			token = tokens[currentIndex]

			// Add elements between quotes as strings
			if token.Value == tokenizer.Quotes {
				ParseStrings(&call)
				continue
			}

			// If we found another parentheses, use recursion
			if tokens[currentIndex+1].Value == tokenizer.LeftParentheses {
				currentIndex++

				call.Params = append(call.Params, Iterate())
				return call
			}

			// Add the element to the params
			call.Params = append(call.Params, Node{
				Type:  token.Type,
				Value: token.Value,
			})
		}
		currentIndex++

		return call
	}

	// Add semicolons tokens
	if token.Type == tokenizer.SemiColonTag {
		currentIndex++
		return Node{
			Type:  SemiColon,
			Value: token.Value,
		}
	}

	currentIndex++

	return Node{}
}

// ParseStrings gets the tokens inside quotes and adds it to the call params
func ParseStrings(call *Node) {
	currentIndex++
	token := tokens[currentIndex]
	var values []string

	// Wait for the other quote to come
	for token.Value != tokenizer.Quotes {
		values = append(values, token.Value)
		currentIndex++
		token = tokens[currentIndex]
	}

	call.Params = append(call.Params, Node{
		Type:  String,
		Value: strings.Join(values, " "),
	})
}
