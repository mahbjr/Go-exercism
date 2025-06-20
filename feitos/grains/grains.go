package grains

import (
	"errors"
)

// Square calculates the number of grains on a given square of a chess board,
// where the number of grains doubles on each square.
// Returns an error if the square is invalid (less than 1 or greater than 64).
func Square(n int) (uint64, error) {
	// Check if the input is valid (1 to 64)
	if n <= 0 || n > 64 {
		return 0, errors.New("square must be between 1 and 64")
	}

	// The number of grains on square n is 2^(n-1)
	// Starting with 1 grain on the first square
	return 1 << (n - 1), nil
}

// Total returns the total number of grains on the chessboard (sum of all squares)
func Total() uint64 {
	// The sum of powers of 2 from 2^0 to 2^63 is 2^64 - 1
	// However, since we're using uint64, we need to handle this calculation carefully
	// to avoid overflow. The result is 2^64 - 1.

	// The elegant mathematical solution is:
	return (1 << 64) - 1

	// Alternative implementation using a loop:
	/*
		var total uint64
		for i := 1; i <= 64; i++ {
			grains, _ := Square(i)
			total += grains
		}
		return total
	*/
}
