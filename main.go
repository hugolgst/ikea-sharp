package main

import "./tokenizer"

func main() {
	example := "TÄRNÖ BEHÅLLAREThis message will be printedBEHÅLLARE"
	tokenizer.Tokenize(example)
}
