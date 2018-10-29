package letter

type FreqMap map[rune]int

// ConcurrentFrequency gets the total number of occurances for each letter in each array
func ConcurrentFrequency(input []string) (output FreqMap) {

	results := make(chan FreqMap)

	for _, value := range input {
		go func(phrase string) {
			results <- Frequency(phrase)
		}(value)
	}

	for i := 0; i < len(input); i++ {
		for key, value := range <-results {
			output[key] += value
		}
	}

	return
}

// Frequency counts how many times a letter appears in a string
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}
