package strand

import "strings"

// ToRNA converts a DNA sequence to its RNA complement.
// The RNA complement is obtained by replacing each nucleotide with its complement:
// G -> C
// C -> G
// T -> A
// A -> U
func ToRNA(dna string) string {
	var result strings.Builder
	result.Grow(len(dna)) // Pre-allocate capacity for efficiency

	for _, nucleotide := range dna {
		switch nucleotide {
		case 'G':
			result.WriteRune('C')
		case 'C':
			result.WriteRune('G')
		case 'T':
			result.WriteRune('A')
		case 'A':
			result.WriteRune('U')
		}
	}

	return result.String()
}
