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

		// Add tokens for the left and right parentheses
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

		// If the character is a number, we iterate all the next characters to see if they are
		// part of the number too.
		if IsNumber(character) {
			var value string

			// Iterate through the next numbers
			for IsNumber(character) {
				value += character
				currentIndex++
				character = string(code[currentIndex])
			}

			// Then add the number token
			tokens = append(tokens, Token{
				Type: "NUMBER",
				Value: value,
			})
		}

		// If the character is a letter, we iterate all the next characters to see if they are
		// part of the letter too.
		if IsLetter(character) {
			var value string

			// Iterate through the next letters
			for IsLetter(character) {
				value += character
				currentIndex++
				character = string(code[currentIndex])
			}

			// Then add the name token
			tokens = append(tokens, Token{
				Type: "NAME",
				Value: value,
			})
		}
	}

	return
}

