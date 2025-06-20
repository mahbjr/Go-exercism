package wordy

import (
	"regexp"
	"strconv"
	"strings"
)

// Answer parses a word problem and returns the calculated result
// The second return value indicates if the calculation was successful (true) or had an error (false)
func Answer(question string) (int, bool) {
	// Clean up the question
	question = strings.TrimSuffix(question, "?")

	// Check if the question starts with "What is"
	if !strings.HasPrefix(question, "What is") {
		return 0, false
	}

	// Remove "What is" from the beginning
	expression := strings.TrimPrefix(question, "What is")
	expression = strings.TrimSpace(expression)

	// If nothing left, it's an invalid question
	if expression == "" {
		return 0, false
	}

	// Use a more comprehensive tokenization approach to catch all words and numbers
	reNumbers := regexp.MustCompile(`(-?\d+)`)
	numbers := reNumbers.FindAllString(expression, -1)

	// Check for patterns that would indicate invalid operations or syntax
	reInvalidOps := regexp.MustCompile(`(\w+)`)
	allWords := reInvalidOps.FindAllString(expression, -1)

	// Validate that we only have supported operations
	validOps := map[string]bool{
		"plus":      true,
		"minus":     true,
		"multiplied": true,
		"by":        true,
		"divided":   true,
	}

	for _, word := range allWords {
		// Skip numbers
		if _, err := strconv.Atoi(word); err == nil {
			continue
		}

		// If it's not a valid operation word, it's an error
		if !validOps[word] {
			return 0, false
		}
	}

	// Now for actual parsing
	re := regexp.MustCompile(`(-?\d+)|plus|minus|(multiplied by)|(divided by)`)
	tokens := re.FindAllString(expression, -1)

	// Check for invalid syntax or empty expression
	if len(tokens) == 0 {
		return 0, false
	}

	// Validate that we have the right number of tokens
	// If the number of numbers doesn't match what we expect based on operations, it's an error
	expectedNumbers := 1 // Start with one number
	for _, token := range tokens {
		if token == "plus" || token == "minus" || token == "multiplied by" || token == "divided by" {
			expectedNumbers++
		}
	}

	if len(numbers) != expectedNumbers {
		return 0, false
	}

	// The first token must be a number
	result, err := strconv.Atoi(tokens[0])
	if err != nil {
		return 0, false
	}

	// Process the remaining tokens
	for i := 1; i < len(tokens); i++ {
		// Each operation must be followed by a number
		switch tokens[i] {
		case "plus":
			if i+1 >= len(tokens) {
				return 0, false
			}
			num, err := strconv.Atoi(tokens[i+1])
			if err != nil {
				return 0, false
			}
			result += num
			i++
		case "minus":
			if i+1 >= len(tokens) {
				return 0, false
			}
			num, err := strconv.Atoi(tokens[i+1])
			if err != nil {
				return 0, false
			}
			result -= num
			i++
		case "multiplied by":
			if i+1 >= len(tokens) {
				return 0, false
			}
			num, err := strconv.Atoi(tokens[i+1])
			if err != nil {
				return 0, false
			}
			result *= num
			i++
		case "divided by":
			if i+1 >= len(tokens) {
				return 0, false
			}
			num, err := strconv.Atoi(tokens[i+1])
			if err != nil {
				return 0, false
			}
			result /= num
			i++
		default:
			// If we get a number when we expect an operator, that's an error
			// This covers "two numbers in a row" case
			_, err := strconv.Atoi(tokens[i])
			if err == nil {
				return 0, false
			}
			// Otherwise, it's an unknown operation
			return 0, false
		}
	}

	return result, true
}
