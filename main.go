package main

import (
	"./parser"
	"./tokenizer"
	"fmt"
)

func main() {
	example := "INNERSKÄR FUNKÖN SKOGSFIBBLA Hello world! SKOGSFIBBLA ÄPPLARÖ FJÄLLBO"
	tokens := tokenizer.Tokenize(example)
	fmt.Println(tokens)
	fmt.Println(parser.Parse(tokens))
}
