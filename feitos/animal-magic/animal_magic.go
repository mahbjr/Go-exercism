package chance

import (
	"math/rand"
	"time"
)

// init initializes the random number generator with a time-based seed
func init() {
	rand.Seed(time.Now().UnixNano())
}

// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {
	return rand.Intn(20) + 1
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
func GenerateWandEnergy() float64 {
	return rand.Float64() * 12.0
}

// ShuffleAnimals returns a slice with all eight animal strings in random order.
func ShuffleAnimals() []string {
	animals := []string{"ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog"}

	// Fisher-Yates shuffle algorithm
	for i := len(animals) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		animals[i], animals[j] = animals[j], animals[i]
	}

	return animals
}
