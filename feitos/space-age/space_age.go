package space

// Planet represents a planet name as a string
type Planet string

// Age calculates the age in years on the specified planet
// Returns -1.0 for invalid planets
func Age(seconds float64, planet Planet) float64 {
	// Earth orbital period in seconds (365.25 days)
	const earthYear float64 = 31557600

	// Orbital periods of planets relative to Earth years
	relativeOrbitalPeriods := map[Planet]float64{
		"Mercury": 0.2408467,
		"Venus":   0.61519726,
		"Earth":   1.0,
		"Mars":    1.8808158,
		"Jupiter": 11.862615,
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 164.79132,
	}

	// Check if the planet is valid
	period, exists := relativeOrbitalPeriods[planet]
	if !exists {
		return -1.0 // Return -1 for invalid planets
	}

	// Calculate the age in Earth years
	earthAge := seconds / earthYear

	// Convert Earth years to the specified planet's years
	return earthAge / period
}
