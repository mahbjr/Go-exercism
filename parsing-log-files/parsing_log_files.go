package parsinglogfiles

import (
	"regexp"
)

// IsValidLine determines if a line is valid based on specified rules.
// A valid line must start with "[ERR]" or "[INF]".
func IsValidLine(text string) bool {
	// Check if the text is too short
	if len(text) < 5 {
		return false
	}

	// Use regexp to check if the text starts with [ERR] or [INF]
	validStart := regexp.MustCompile(`^\[(ERR|INF)\]`)
	return validStart.MatchString(text)
}

// SplitLogLine splits a log line into sections based on the delimiter pattern.
// The delimiter is anything between angular brackets (e.g., <*>, <~~~>, etc.)
func SplitLogLine(text string) []string {
	// Create a regular expression to match any text between angle brackets
	// including empty angle brackets <>
	re := regexp.MustCompile(`<[^>]*>`)

	// Split the text using the regex as delimiter
	return re.Split(text, -1)
}

// CountQuotedPasswords counts occurrences of the word "password" (case insensitive)
// within quotes in a set of log lines.
func CountQuotedPasswords(lines []string) int {
	count := 0
	re := regexp.MustCompile(`(?i)"[^"]*password[^"]*"`)

	for _, line := range lines {
		// Count matches in the current line
		matches := re.FindAllString(line, -1)
		count += len(matches)
	}

	return count
}

// RemoveEndOfLineText removes all instances of text following the pattern "end-of-line"
// followed by digits.
func RemoveEndOfLineText(text string) string {
	// Create a regular expression to match "end-of-line" followed by digits
	re := regexp.MustCompile(`end-of-line\d+`)

	// Replace all matches with an empty string
	return re.ReplaceAllString(text, "")
}

// TagWithUserName adds a user tag to the beginning of lines that contain
// "User" followed by a username.
func TagWithUserName(lines []string) []string {
	result := make([]string, len(lines))
	// Regex to find username after "User"
	re := regexp.MustCompile(`User\s+(\w+)`)

	for i, line := range lines {
		// Find the first username in the line
		matches := re.FindStringSubmatch(line)
		if len(matches) > 1 {
			// Add user tag to the beginning of the line
			username := matches[1]
			result[i] = "[USR] " + username + " " + line
		} else {
			// Keep the line unchanged
			result[i] = line
		}
	}

	return result
}
