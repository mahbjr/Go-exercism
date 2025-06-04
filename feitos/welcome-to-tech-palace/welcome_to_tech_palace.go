package techpalace

import (
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	// Create a border line with the specified number of stars
	border := strings.Repeat("*", numStarsPerLine)

	// Combine the border and message
	return border + "\n" + welcomeMsg + "\n" + border
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	// Remove all asterisks
	noStars := strings.ReplaceAll(oldMsg, "*", "")

	// Split by newlines and join with a space
	lines := strings.Split(noStars, "\n")

	// Trim each line and join them together
	var cleanLines []string
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed != "" {
			cleanLines = append(cleanLines, trimmed)
		}
	}

	// Join the clean lines with spaces
	result := strings.Join(cleanLines, " ")

	// Trim any remaining whitespace
	return strings.TrimSpace(result)
}
