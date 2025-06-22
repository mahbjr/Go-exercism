package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

// Encode implements the classic crypto square code
func Encode(input string) string {
	// Handle empty input
	if input == "" {
		return ""
	}

	// Normalize the input (remove non-alphanumeric, convert to lowercase)
	normalized := normalize(input)

	// Handle empty normalized input
	if len(normalized) == 0 {
		return ""
	}

	// Calculate rectangle dimensions
	c := int(math.Ceil(math.Sqrt(float64(len(normalized)))))
	r := int(math.Ceil(float64(len(normalized)) / float64(c)))

	// Pad the normalized string to fill the rectangle
	paddedText := normalized
	for len(paddedText) < r*c {
		paddedText += " "
	}

	// Organize into a grid
	grid := make([][]rune, r)
	for i := range grid {
		grid[i] = make([]rune, c)
		for j := 0; j < c; j++ {
			idx := i*c + j
			if idx < len(paddedText) {
				grid[i][j] = rune(paddedText[idx])
			} else {
				grid[i][j] = ' '
			}
		}
	}

	// Read columns and build the encoded result
	var result strings.Builder
	for j := 0; j < c; j++ {
		if j > 0 {
			result.WriteRune(' ')
		}
		for i := 0; i < r; i++ {
			result.WriteRune(grid[i][j])
		}
	}

	return result.String()
}

// normalize removes all non-alphanumeric characters and converts to lowercase
func normalize(input string) string {
	var normalized strings.Builder

	for _, char := range input {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			normalized.WriteRune(unicode.ToLower(char))
		}
	}

	return normalized.String()
}
