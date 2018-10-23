package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram returns true if the provided word is an isogram, false otherwise
func IsIsogram(word string) bool {
	word = strings.ToUpper(word)
	usedLetters := make(map[rune]bool, 26)

	for _, value := range word {
		_, used := usedLetters[value]
		if used {
			return false
		}

		if unicode.IsLetter(value) {
			usedLetters[value] = true
		}
	}

	return true
}
