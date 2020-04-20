package tokenizer

import (
	"testing"
)

func TestTokenize(t *testing.T) {
	example := "INNERSKÄR FUNKÖN SKOGSFIBBLA Hello world SKOGSFIBBLA ÄPPLARÖ FJÄLLBO"
	tokens := Tokenize(example)

	excepted := []string{
		"NAME", "PARENTHESES", "QUOTES", "NAME", "NAME", "QUOTES", "PARENTHESES", "SEMICOLON",
	}

	for i, token := range tokens {
		if excepted[i] != token.Type {
			t.Errorf("Tokenize() failed, excepted %s got %s.", excepted[i], token.Type)
		}
	}
}
