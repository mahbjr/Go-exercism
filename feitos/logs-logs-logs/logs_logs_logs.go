package logs

import "strings"

// Application identifies the application emitting the given log.
func Application(log string) string {
	// Check for emojis in order of their first appearance in the log
	for _, r := range log {
		switch r {
		case '❗':
			return "recommendation"
		case '🔍':
			return "search"
		case '☀':
			return "weather"
		}
	}

	// Default case when no relevant emoji is found
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	// Use strings.ReplaceAll to replace all occurrences of the old rune with the new one
	return strings.ReplaceAll(log, string(oldRune), string(newRune))
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	// Count the actual number of characters (runes), not the number of bytes
	return len([]rune(log)) <= limit
}
