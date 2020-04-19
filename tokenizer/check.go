package tokenizer

import "regexp"

// IsNumber checks if the given character is a number and returns the condition
func IsNumber(character string) bool {
	numberRegex := regexp.MustCompile(`\d`)

	return numberRegex.Match([]byte(character))
}

// IsLetter checks if the given character is a letter and returns the condition
func IsLetter(character string) bool {
	letterRegex := regexp.MustCompile(`[a-zA-ZÖÄÅ_]`)

	return letterRegex.Match([]byte(character))
}
