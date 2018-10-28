package luhn

import (
	"strings"
	"unicode"
)

// Valid returns true if the given input is valid, false otherwise
func Valid(input string) bool {

	input = strings.Replace(input, " ", "", -1)

	if len(input) <= 1 {
		return false
	}

	sum := 0
	input = reverseString(input)

	for key, value := range input {

		if unicode.IsDigit(value) == false {
			return false
		}

		number := int(value - '0')

		if key%2 == 0 {
			sum += number
			continue
		}

		if number*2 > 9 {
			number = (number * 2) - 9
		} else {
			number *= 2
		}

		sum += number
	}

	if sum%10 == 0 {
		return true
	}

	return false
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
