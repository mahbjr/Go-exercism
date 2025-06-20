package isbn

import (
	"strings"
	"unicode"
)

// IsValidISBN checks if the given string is a valid ISBN-10 number
func IsValidISBN(isbn string) bool {
	// Remove dashes
	isbn = strings.ReplaceAll(isbn, "-", "")

	// Check length - valid ISBN-10 has 10 digits
	if len(isbn) != 10 {
		return false
	}

	sum := 0
	for i, char := range isbn {
		// Handle the check digit (last position) separately
		if i == 9 && char == 'X' {
			sum += 10 * (10 - i)
			continue
		}

		// Ensure it's a digit
		if !unicode.IsDigit(char) {
			return false
		}

		// Convert to integer and add to weighted sum
		digit := int(char - '0')
		sum += digit * (10 - i)
	}

	// Valid ISBN-10 should be divisible by 11
	return sum%11 == 0
}
