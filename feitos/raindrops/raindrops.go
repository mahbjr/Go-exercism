package raindrops

import "strconv"

// Convert converts a number to a string, the contents of which depend on the number's factors.
func Convert(number int) string {
	var result string

	// Check for factors and append appropriate sounds
	if number%3 == 0 {
		result += "Pling"
	}
	if number%5 == 0 {
		result += "Plang"
	}
	if number%7 == 0 {
		result += "Plong"
	}

	// If no factors matched, return the number as a string
	if result == "" {
		return strconv.Itoa(number)
	}

	return result
}
