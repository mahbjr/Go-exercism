package dna

import "errors"

// Histogram is a mapping from nucleotide to its count in given DNA.
// Choose a suitable data type.
type Histogram map[rune]int

// DNA is a list of nucleotides
type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
func (d DNA) Counts() (Histogram, error) {
	counts := Histogram{
		'A': 0,
		'C': 0,
		'G': 0,
		'T': 0,
	}

	for _, nucleotide := range d {
		// Check if the nucleotide is valid
		if _, valid := counts[nucleotide]; !valid {
			return nil, errors.New("invalid nucleotide in DNA strand")
		}

		// Increment the count for this nucleotide
		counts[nucleotide]++
	}

	return counts, nil
}
