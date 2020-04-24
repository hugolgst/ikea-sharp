package tokenizer

import "regexp"

// IsParentheses checks if the given string is a parentheses or not
func IsParentheses(content string) bool {
	return content == LeftParentheses || content == RightParentheses
}

// IsNumber checks if the given character is a number and returns the condition
func IsNumber(character string) bool {
	numberRegex := regexp.MustCompile(`\d`)

	return numberRegex.Match([]byte(character))
}

// IsLetter checks if the given character is a letter and returns the condition
func IsLetter(character string) bool {
	letterRegex := regexp.MustCompile(`[a-zA-Z]|[à-ú]|[À-Ú]|[!%=]`)

	return letterRegex.Match([]byte(character))
}
