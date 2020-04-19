package tokenizer

// A Token contains the type and the value of a part of code
type Token struct {
	Type  string
	Value string
}

const (
	LeftParentheses  = "("
	RightParentheses = ")"

	Space = " "
)

// Tokenize takes the given code, tokenize it and returns an array of Tokens
func Tokenize(code string) (tokens []Token) {
	var currentIndex int

	// Iterate each character of the code
	for currentIndex < len(code) {
		character := string(code[currentIndex])

		// Add parentheses tokens
		if character == LeftParentheses || character == RightParentheses {
			tokens = append(tokens, Token{
				Type:  "PARENTHESES",
				Value: character,
			})

			currentIndex++
			continue
		}

		// Skip the spaces
		if character == Space {
			currentIndex++
			continue
		}
	}

	return
}

