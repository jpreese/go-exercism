// Package isogram provides helper functions
// for working with isograms and text input
package isogram

import (
	"sort"
	"strings"
	"unicode"
)

// IsIsogramSortWord takes a string of input, and determines if it is an isogram
// (as in doesn't have any repeat letters)
func IsIsogramSortWord(word string) bool {
	runes := []rune(word)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	for i := range runes {
		r := unicode.ToLower(runes[i])
		if r == ' ' || r == '-' {
			continue
		}
		if i < len(runes)-1 && r == unicode.ToLower(runes[i+1]) {
			return false
		}
	}
	return true
}

///////////////////////////////////////////////

// IsIsogramUsingMap takes a string of input, and determines if it is an isogram
// (as in doesn't have any repeat letters)
func IsIsogramUsingMap(word string) bool {
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

////////////////////////////////////////////////

// IsIsogramAppendRune takes a string of input, and determines if it is an isogram
// (as in doesn't have any repeat letters)
func IsIsogramAppendRune(input string) bool {
	charsMet := []rune{}

	for _, r := range input {
		r := unicode.ToLower(r)
		// check if this element is already
		// in charsMet and if so, return false
		if runeInSlice(r, charsMet) {
			// in this case, it's not an isogram
			return false
		}

		// push the rune into the slice
		// if it's not hyphen or space
		if r != ' ' && r != '-' {
			charsMet = append(charsMet, r)
		}
	}

	return true
}

func runeInSlice(r rune, s []rune) bool {
	for _, v := range s {
		if v == r {
			return true
		}
	}
	return false
}

/////////////////////////////////////////////
