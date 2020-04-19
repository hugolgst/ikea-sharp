package tokenizer

import "regexp"

// IsNumber checks if the given character is a number and returns the condition
func IsNumber(character string) bool {
	numberRegex := regexp.MustCompile(`\d`)

	return numberRegex.Match([]byte(character))
}
