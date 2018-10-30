package letter

// FrequencyMap holds a rune and the number of occurances
type FrequencyMap map[rune]int

// ConcurrentFrequency gets the total number of occurances for each letter in each array
func ConcurrentFrequency(input []string) FrequencyMap {

	results := make(chan FrequencyMap)
	output := make(FrequencyMap)

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

	return output
}

// Frequency counts how many times a letter appears in a string
func Frequency(s string) FrequencyMap {
	m := FrequencyMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}
