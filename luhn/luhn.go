package luhn

import (
	"fmt"
	"strings"
)

// Valid returns true if the given input is valid, false otherwise
func Valid(input string) bool {
	if len(input) <= 1 {
		return false
	}

	sum := 0
	input = strings.Replace(input, " ", "", -1)

	fmt.Println("input is " + input)

	for i := 0; i < len(input); i++ {
		fmt.Printf("checking %v\n", input[i])

		if i%2 == 0 {
			//input[i] = input[i] * 2
		}

		fmt.Println(input[i])

		sum += int(input[i])
	}

	if sum%10 == 0 {
		return true
	}

	return false
}
