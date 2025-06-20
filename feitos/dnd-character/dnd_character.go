package dndcharacter

import (
	"math"
	"math/rand"
)

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	// The standard D&D formula for ability modifiers: (score - 10) / 2, rounded down
	return int(math.Floor(float64(score-10) / 2.0))
}

// Character represents a D&D character
type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Ability generates a random ability score (typically by rolling 4d6 and taking the sum of the highest 3)
func Ability() int {
	// Roll 4d6
	rolls := []int{
		rand.Intn(6) + 1,
		rand.Intn(6) + 1,
		rand.Intn(6) + 1,
		rand.Intn(6) + 1,
	}

	// Find the minimum roll (to be discarded)
	minRoll := rolls[0]
	minIndex := 0
	for i, roll := range rolls {
		if roll < minRoll {
			minRoll = roll
			minIndex = i
		}
	}

	// Sum the three highest dice rolls
	sum := 0
	for i, roll := range rolls {
		if i != minIndex {
			sum += roll
		}
	}

	return sum
}

// GenerateCharacter creates a new Character with random ability scores
func GenerateCharacter() Character {
	// Generate random ability scores
	constitution := Ability()

	return Character{
		Strength:     Ability(),
		Dexterity:    Ability(),
		Constitution: constitution,
		Intelligence: Ability(),
		Wisdom:       Ability(),
		Charisma:     Ability(),
		// Hitpoints = 10 + constitution modifier
		Hitpoints:    10 + Modifier(constitution),
	}
}
