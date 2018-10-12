package strand

import "strings"

// ToRNA converts a given DNA strand to an RNA strand
func ToRNA(dna string) string {

	complements := map[byte]byte{
		'G': 'C',
		'C': 'G',
		'T': 'A',
		'A': 'U',
	}

	var rna strings.Builder
	for i := 0; i < len(dna); i++ {
		rna.WriteByte(complements[dna[i]])
	}

	return rna.String()
}
