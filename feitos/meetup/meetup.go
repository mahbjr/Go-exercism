package meetup

import "time"

// WeekSchedule represents different schedules for meetups
type WeekSchedule int

const (
	First  WeekSchedule = 1
	Second WeekSchedule = 2
	Third  WeekSchedule = 3
	Fourth WeekSchedule = 4
	Teenth WeekSchedule = 5
	Last   WeekSchedule = 6
)

// Day returns the day of the month for a given meetup.
func Day(week WeekSchedule, weekday time.Weekday, month time.Month, year int) int {
	switch week {
	case First, Second, Third, Fourth:
		// First day of the month
		startDate := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)

		// Find the first occurrence of the weekday
		daysToAdd := (int(weekday) - int(startDate.Weekday()) + 7) % 7
		firstOccurrence := time.Date(year, month, 1+daysToAdd, 0, 0, 0, 0, time.UTC)

		// Add the necessary weeks to get to the desired week (first, second, third, or fourth)
		meetupDate := firstOccurrence.AddDate(0, 0, int(week-1)*7)
		return meetupDate.Day()

	case Teenth:
		// Start from the 13th of the month
		teenthStart := time.Date(year, month, 13, 0, 0, 0, 0, time.UTC)

		// Find the next occurrence of the weekday after the 13th
		daysToAdd := (int(weekday) - int(teenthStart.Weekday()) + 7) % 7
		meetupDate := teenthStart.AddDate(0, 0, daysToAdd)
		return meetupDate.Day()

	case Last:
		// Get the first day of the next month, then go back one day to get the last day of the current month
		lastDay := time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC)

		// Find the last occurrence of the weekday by going backwards from the last day
		daysToSubtract := (int(lastDay.Weekday()) - int(weekday) + 7) % 7
		meetupDate := lastDay.AddDate(0, 0, -daysToSubtract)
		return meetupDate.Day()
	}

	return 0 // Should never reach here
}
