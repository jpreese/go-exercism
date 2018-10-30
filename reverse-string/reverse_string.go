package reverse

// String takes a string as input and reverses the string
func String(input string) string {
	runes := []rune(input)

	for x, y := 0, len(runes)-1; x < y; x, y = x+1, y-1 {
		runes[x], runes[y] = runes[y], runes[x]
	}

	return string(runes)
}
