package proverb

import (
	"fmt"
)

// Given an array of proverb inputs return the corresponding proverb lines
func Proverb(rhyme []string) []string {

	result := make([]string, len(rhyme))

	if len(rhyme) == 0 {
		return result
	}

	if len(rhyme) == 1 {
		result[0] = getLastProverbLine(rhyme[0])
		return result
	}

	for i := 0; i < len(rhyme)-1; i++ {
		result[i] = fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1])
	}

	result[len(result)-1] = getLastProverbLine(rhyme[0])
	return result
}

func getLastProverbLine(input string) string {
	return fmt.Sprintf("And all for the want of a %s.", input)
}
