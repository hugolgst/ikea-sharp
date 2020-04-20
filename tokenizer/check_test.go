package tokenizer

import (
	"testing"
)

func TestIsNumber(t *testing.T) {
	number := "1"

	if !IsNumber(number) {
		t.Errorf("IsNumber() failed.")
	}
}

func TestIsLetter(t *testing.T) {
	sentence := "INNERSKÄRFUNKÖNSKOGSFIBBLA"

	for _, letter := range sentence {
		if !IsLetter(string(letter)) {
			t.Errorf("IsLetter() failed, excepted %s to be a letter.", string(letter))
		}
	}
}
