package luhn

import (
	"strings"
	"unicode"
)

// Valid determines whether a string is valid according to the Luhn algorithm
func Valid(id string) bool {
	// Remove spaces
	id = strings.ReplaceAll(id, " ", "")

	// Check if the input has at least 2 digits
	if len(id) <= 1 {
		return false
	}

	// Check if the input contains only digits
	for _, c := range id {
		if !unicode.IsDigit(c) {
			return false
		}
	}

	// Apply the Luhn algorithm
	sum := 0
	for i := len(id) - 1; i >= 0; i-- {
		digit := int(id[i] - '0')

		// Double every second digit from right to left
		if (len(id)-i)%2 == 0 {
			digit *= 2
			// If doubling results in a number greater than 9, subtract 9
			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
	}

	// If the sum is divisible by 10, the number is valid
	return sum%10 == 0
}
