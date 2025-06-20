// Package leap provides functionality for determining leap years.
package leap

// IsLeapYear determines whether a given year is a leap year.
// The rules for leap years are:
// - Years divisible by 4 are leap years
// - But years divisible by 100 are not leap years
// - Unless they are also divisible by 400, in which case they are leap years
func IsLeapYear(year int) bool {
	// divisible by 400: leap year
	if year%400 == 0 {
		return true
	}

	// divisible by 100: not a leap year
	if year%100 == 0 {
		return false
	}

	// divisible by 4: leap year
	if year%4 == 0 {
		return true
	}

	// not divisible by 4: not a leap year
	return false
}
