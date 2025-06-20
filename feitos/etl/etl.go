package etl

import "strings"

// Transform converts a legacy data format (score -> letters) to a new format (letter -> score)
func Transform(input map[int][]string) map[string]int {
	result := make(map[string]int)

	for score, letters := range input {
		for _, letter := range letters {
			// Convert letter to lowercase and add to the result map with its score
			result[strings.ToLower(letter)] = score
		}
	}

	return result
}
