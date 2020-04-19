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
