package main

import (
	"./parser"
	"./runtime"
	"./tokenizer"
)

func main() {
	example := `
		TILLGÅNG FUNKÖN SKOGSFIBBLA hey SKOGSFIBBLA SKOGSFIBBLA nik zebi SKOGSFIBBLA ÄPPLARÖ FJÄLLBO
		SMÅGLI FUNKÖN SMÅKALLT FUNKÖN SKOGSFIBBLA hey SKOGSFIBBLA ÄPPLARÖ ÄPPLARÖ FJÄLLBO
	`

	tokens := tokenizer.Tokenize(example)
	ast := parser.Parse(tokens)
	runtime.Run(ast)
}
