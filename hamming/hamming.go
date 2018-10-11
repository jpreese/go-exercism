// Provides functionality to calculate the Hamming
// distance between two DNA strands
package hamming

import "errors"

// Calculates the Hamming distance between two
// DNA strands
func Distance(a, b string) (int, error) {

	if len(a) != len(b) {
		return -1, errors.New("size of both strings must match")
	}

	distance := 0

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}

	return distance, nil
}
