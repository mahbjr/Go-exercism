package rotationalcipher

// RotationalCipher encodes a string by rotating each letter by shiftKey positions
func RotationalCipher(plain string, shiftKey int) string {
	// Normalize the shift key to be between 0 and 25
	shiftKey = shiftKey % 26

	// If no shift, return the original string
	if shiftKey == 0 {
		return plain
	}

	// Create a byte array to store the result
	result := make([]byte, len(plain))

	for i := 0; i < len(plain); i++ {
		char := plain[i]

		switch {
		// Uppercase letters: A-Z (ASCII 65-90)
		case 'A' <= char && char <= 'Z':
			// Rotate and handle wrap-around
			result[i] = 'A' + byte((int(char-'A')+shiftKey)%26)

		// Lowercase letters: a-z (ASCII 97-122)
		case 'a' <= char && char <= 'z':
			// Rotate and handle wrap-around
			result[i] = 'a' + byte((int(char-'a')+shiftKey)%26)

		// Non-letters: keep unchanged
		default:
			result[i] = char
		}
	}

	return string(result)
}
