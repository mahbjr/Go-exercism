package reverse

// Reverse returns the input string with its characters in reverse order
func Reverse(input string) string {
	// Convert string to a slice of runes to handle Unicode characters properly
	runes := []rune(input)

	// Get the length of the rune slice
	length := len(runes)

	// Create a new slice of runes with the same length
	reversed := make([]rune, length)

	// Fill the new slice in reverse order
	for i, char := range runes {
		reversed[length-i-1] = char
	}

	// Convert back to string and return
	return string(reversed)
}
