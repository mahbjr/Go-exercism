package bookstore

import (
	"sort"
)

// Cost calculates the total cost of a basket of books with applicable discounts
func Cost(books []int) int {
	if len(books) == 0 {
		return 0
	}

	// Count occurrences of each book
	counts := make(map[int]int)
	for _, book := range books {
		counts[book]++
	}

	// Extract the counts
	bookCounts := make([]int, 0, len(counts))
	for _, count := range counts {
		bookCounts = append(bookCounts, count)
	}

	// Sort counts in descending order to optimize grouping
	sort.Sort(sort.Reverse(sort.IntSlice(bookCounts)))

	// Calculate the best price using dynamic approach
	return findBestPrice(bookCounts)
}

// findBestPrice calculates the minimum price for the given book counts
func findBestPrice(bookCounts []int) int {
	// Base case: no books
	if allZero(bookCounts) {
		return 0
	}

	// Try all possible group sizes and pick the best one
	minPrice := 1<<31 - 1 // Max int

	// Try purchasing 1, 2, 3, 4, or 5 different books
	for groupSize := 1; groupSize <= 5 && groupSize <= len(bookCounts); groupSize++ {
		// Create a copy of book counts
		newCounts := make([]int, len(bookCounts))
		copy(newCounts, bookCounts)

		// Take one of each book up to groupSize
		nonZeroBooks := 0
		for i := 0; i < len(newCounts) && nonZeroBooks < groupSize; i++ {
			if newCounts[i] > 0 {
				newCounts[i]--
				nonZeroBooks++
			}
		}

		// If we couldn't form a group of the current size, skip
		if nonZeroBooks < groupSize {
			continue
		}

		// Remove any books with zero counts to optimize
		var filteredCounts []int
		for _, count := range newCounts {
			if count > 0 {
				filteredCounts = append(filteredCounts, count)
			}
		}

		// Calculate price for this group + best price for remaining books
		groupPrice := calculateGroupPrice(groupSize)
		remainderPrice := findBestPrice(filteredCounts)
		totalPrice := groupPrice + remainderPrice

		if totalPrice < minPrice {
			minPrice = totalPrice
		}
	}

	return minPrice
}

// calculateGroupPrice returns the price for a group of different books
func calculateGroupPrice(groupSize int) int {
	basePrice := 800 // Price per book in cents
	discounts := map[int]float64{
		1: 0.00,
		2: 0.05,
		3: 0.10,
		4: 0.20,
		5: 0.25,
	}

	return int(float64(basePrice*groupSize) * (1 - discounts[groupSize]))
}

// allZero checks if all counts are zero
func allZero(counts []int) bool {
	for _, count := range counts {
		if count > 0 {
			return false
		}
	}
	return true
}
