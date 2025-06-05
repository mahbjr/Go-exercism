package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	total := 0
	for _, count := range birdsPerDay {
		total += count
	}
	return total
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	// Each week has 7 days
	startIndex := (week - 1) * 7
	endIndex := startIndex + 7

	// Make sure we don't go out of bounds
	if endIndex > len(birdsPerDay) {
		endIndex = len(birdsPerDay)
	}

	total := 0
	for i := startIndex; i < endIndex; i++ {
		total += birdsPerDay[i]
	}

	return total
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	// For each even-indexed day (0, 2, 4, ...), increment the bird count by 1
	for i := 0; i < len(birdsPerDay); i += 2 {
		birdsPerDay[i]++
	}

	return birdsPerDay
}
