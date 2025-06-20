// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
	"unicode"
)

// Hey returns Bob's response to a given remark
func Hey(remark string) string {
	// Trim whitespace
	remark = strings.TrimSpace(remark)

	// Check if silent
	if remark == "" {
		return "Fine. Be that way!"
	}

	// Check if it's a question (ends with ?)
	isQuestion := strings.HasSuffix(remark, "?")

	// Check if shouting (has at least one letter and all letters are uppercase)
	hasLetters := false
	allCaps := true

	for _, char := range remark {
		if unicode.IsLetter(char) {
			hasLetters = true
			if !unicode.IsUpper(char) {
				allCaps = false
				break
			}
		}
	}

	isShouting := hasLetters && allCaps

	// Determine response based on combination of factors
	switch {
	case isShouting && isQuestion:
		return "Calm down, I know what I'm doing!"
	case isShouting:
		return "Whoa, chill out!"
	case isQuestion:
		return "Sure."
	default:
		return "Whatever."
	}
}
