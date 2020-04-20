package tokenizer

import (
	"fmt"
	"testing"
)

func TestTokenize(t *testing.T) {
	example := "INNERSKÄR FUNKÖN 213 ÄPPLARÖ FJÄLLBO"
	tokens := Tokenize(example)

	fmt.Println(tokens)
}
