// Package hamming provides functionality to calculate the Hamming
// distance between two DNA strands
package hamming

import "errors"

// Distance calculates the Hamming distance between two
// DNA strands
func Distance(a, b string) (int, error) {

	if len(a) != len(b) {
		return -1, errors.New("size of both strings must match")
	}

	distance := 0

	for key := range a {
		if a[key] != b[key] {
			distance++
		}
	}

	return distance, nil
}
