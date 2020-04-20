package main

import (
	"./log"
	"./parser"
	"./runtime"
	"./tokenizer"
	"io/ioutil"
	"os"
)

func main() {
	// Check if the file is specified
	if len(os.Args) < 2 {
		log.Errorf("You need to give the IKEA# file to execute.")
	}

	// Read the given file
	bytes, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Errorf(err)
	}

	// Execute the code
	tokens := tokenizer.Tokenize(string(bytes))
	ast := parser.Parse(tokens)
	runtime.Run(ast)
}
