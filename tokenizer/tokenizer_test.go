package tokenizer

import (
	"testing"
)

func TestTokenize(t *testing.T) {
	example := "hello (124)"
	tokens := Tokenize(example)

	excepted := []string{
		"NAME", "PARENTHESES", "NUMBER", "PARENTHESES",
	}

	for i, token := range tokens {
		if token.Type != excepted[i] {
			t.Errorf("Tokenize() failed, excepted %s got %s.", excepted[i], token.Type)
		}
	}
}
