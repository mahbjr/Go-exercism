package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram determines whether a word is an isogram.
// An isogram is a word or phrase without a repeating letter,
// ignoring case and non-letter characters.
func IsIsogram(word string) bool {
	// Convert to lowercase to make the check case-insensitive
	word = strings.ToLower(word)

	// Use a map to track which letters we've seen
	seen := make(map[rune]bool)

	for _, char := range word {
		// Only check letters, ignore spaces, hyphens, etc.
		if unicode.IsLetter(char) {
			// If we've seen this letter before, it's not an isogram
			if seen[char] {
				return false
			}
			// Mark this letter as seen
			seen[char] = true
		}
	}

	// If we get here, no letter appears more than once
	return true
}
