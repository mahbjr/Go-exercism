package darts

import (
	"math"
)

// Score calculates the score for a dart throw based on the distance from the center
func Score(x, y float64) int {
	// Calculate the distance from the center (0,0) using the Pythagorean theorem
	distance := math.Sqrt(x*x + y*y)

	// Determine the score based on which circle the dart landed in
	switch {
	case distance <= 1.0:
		// Inner circle (radius 1) - 10 points
		return 10
	case distance <= 5.0:
		// Middle circle (radius 5) - 5 points
		return 5
	case distance <= 10.0:
		// Outer circle (radius 10) - 1 point
		return 1
	default:
		// Outside all circles - 0 points
		return 0
	}
}
