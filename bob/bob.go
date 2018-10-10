package bob

import "strings"
import "unicode"

func Hey(remark string) string {

	remark = strings.TrimSpace(remark)

	if isQuiet(remark) {
		return "Fine. Be that way!"
	}

	if isYelling(remark) && isQuestion(remark) {
		return "Calm down, I know what I'm doing!"
	}

	if isYelling(remark) {
		return "Whoa, chill out!"
	}

	if isQuestion(remark) {
		return "Sure."
	}

	return "Whatever."
}

func isQuiet(remark string) bool {
	return remark == ""
}

func isYelling(remark string) bool {
	return strings.ToUpper(remark) == remark && containsLetter(remark)
}

func isQuestion(remark string) bool {
	return strings.HasSuffix(remark, "?")
}

func containsLetter(input string) bool {
	for _, val := range input {
		if unicode.IsLetter(val) {
			return true
		}
	}

	return false
}
