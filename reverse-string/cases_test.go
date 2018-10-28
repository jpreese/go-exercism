package reverse

// Source: exercism/problem-specifications
// Commit: 2f77985 reverse-string: apply "input" policy
// Problem Specifications Version: 1.1.0

type reverseTestCase struct {
	description string
	input       string
	expected    string
}

var testCases = []reverseTestCase{

	{
		description: "a palindrome",
		input:       "racecar",
		expected:    "racecar",
	},
}
