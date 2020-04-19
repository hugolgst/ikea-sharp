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
	letter := "a"

	if !IsLetter(letter) {
		t.Errorf("IsLetter() failed.")
	}
}
