package hamming

import "errors"

// Distance calculates the Hamming Distance between two DNA strands.
// It returns an error if the strands are of different lengths.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("strands must be of equal length")
	}

	distance := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}

	return distance, nil
}
