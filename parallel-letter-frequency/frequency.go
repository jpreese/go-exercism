package letter

type FreqMap map[rune]int

// ConcurrentFrequency gets the total number of occurances for each letter in each array
func ConcurrentFrequency(input []string) FreqMap {

	results := make(chan FreqMap)
	freqMap := make(FreqMap)

	for i := 0; i < len(input); i++ {
		go worker(input[i], results)
	}

	for j := 0; j < len(input); j++ {
		addToMap(freqMap, <-results)
	}

	return freqMap
}

// Frequency counts how many times a letter appears in a string
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func worker(phrase string, results chan<- FreqMap) {
	results <- Frequency(phrase)
}

func addToMap(finalResult FreqMap, partialResult FreqMap) {
	for key, value := range partialResult {
		_, exists := finalResult[key]
		if exists {
			finalResult[key] = finalResult[key] + value
		} else {
			finalResult[key] = partialResult[key]
		}
	}
}
