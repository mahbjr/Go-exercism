package pangram

import (
	"strings"
	"unicode"
)

// IsPangram determines if a string is a pangram.
// A pangram is a sentence that contains every letter of the alphabet at least once.
func IsPangram(input string) bool {
	// Convert to lowercase for case-insensitive check
	input = strings.ToLower(input)

	// Use a map to track which letters we've seen
	letters := make(map[rune]bool)

	// Scan through the string
	for _, char := range input {
		// Only count letters
		if unicode.IsLetter(char) {
			letters[char] = true
		}
	}

	// Check if we've seen all 26 letters of the English alphabet
	return len(letters) == 26
}
