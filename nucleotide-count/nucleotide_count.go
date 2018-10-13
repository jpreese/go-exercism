package dna

import "errors"

// Histogram represents the frequency of each DNA symbol
type Histogram map[rune]int

// DNA is a string containing a DNA sequence
type DNA string

// Counts returns a Histrogram of the current DNA sequence
func (d DNA) Counts() (Histogram, error) {

	result := map[rune]int{
		'A': 0,
		'T': 0,
		'C': 0,
		'G': 0,
	}

	for _, val := range d {

		_, exists := result[val]
		if exists == false {
			return nil, errors.New("invalid nucleotide found")
		}

		result[val]++
	}

	return result, nil
}
