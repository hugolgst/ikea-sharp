package main

import (
	"./parser"
	"./tokenizer"
	"fmt"
)

func main() {
	example := "INNERSKÄR FUNKÖN INNERSKÄR FUNKÖN 124 ÄPPLARÖ ÄPPLARÖ FJÄLLBO"
	tokens := tokenizer.Tokenize(example)
	fmt.Println(tokens)
	fmt.Println(parser.Parse(tokens))
}
