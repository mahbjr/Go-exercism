package collatzconjecture

import "errors"

// CollatzConjecture returns the number of steps required to reach 1
// starting from the given positive integer n
func CollatzConjecture(n int) (int, error) {
	// Check for invalid input - the conjecture is only defined for positive integers
	if n <= 0 {
		return 0, errors.New("input must be a positive integer")
	}

	// Count the steps
	steps := 0

	// Continue until we reach 1
	for n != 1 {
		if n%2 == 0 {
			// If n is even, divide by 2
			n = n / 2
		} else {
			// If n is odd, multiply by 3 and add 1
			n = 3*n + 1
		}
		steps++
	}

	return steps, nil
}
