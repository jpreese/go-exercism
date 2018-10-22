package acronym

import "strings"

// Abbreviate abbreviates a string into an acronym
func Abbreviate(s string) string {
	var result string
	words := strings.FieldsFunc(s, split)

	for _, value := range words {
		result += string(value[0])
	}

	return strings.ToUpper(result)
}

func split(r rune) bool {
	return r == ' ' || r == '-'
}
