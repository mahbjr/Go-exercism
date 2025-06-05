package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	// Parse the date string in the format MM/DD/YYYY HH:MM:SS
	t, err := time.Parse("1/2/2006 15:04:05", date)
	if err != nil {
		// This shouldn't happen with valid inputs from the test
		panic(fmt.Sprintf("Failed to parse date: %v", err))
	}
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	// Parse the date string in the format "Month D, YYYY HH:MM:SS"
	t, err := time.Parse("January 2, 2006 15:04:05", date)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse date: %v", err))
	}

	// Compare with current time
	return time.Now().After(t)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	// Parse the date string in the format "Day, Month D, YYYY HH:MM:SS"
	t, err := time.Parse("Monday, January 2, 2006 15:04:05", date)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse date: %v", err))
	}

	// Check if hour is between 12 and 18 (noon to 6 PM)
	hour := t.Hour()
	return hour >= 12 && hour < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	// Parse the date string in the format MM/DD/YYYY HH:MM:SS
	t, err := time.Parse("1/2/2006 15:04:05", date)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse date: %v", err))
	}

	// Format the date according to the specified output format
	formattedDate := t.Format("Monday, January 2, 2006, at 15:04.")
	return "You have an appointment on " + formattedDate
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	// Return September 15th of the current year
	currentYear := time.Now().Year()
	return time.Date(currentYear, time.September, 15, 0, 0, 0, 0, time.UTC)
}
