package thefarm

import (
	"errors"
	"fmt"
)

// DivideFood calculates how much fodder each cow gets based on the total fodder amount
// and fattening factor.
func DivideFood(calc FodderCalculator, cows int) (float64, error) {
	// Get the total fodder amount
	fodderAmount, err := calc.FodderAmount(cows)
	if err != nil {
		return 0, err
	}

	// Get the fattening factor
	fatteningFactor, err := calc.FatteningFactor()
	if err != nil {
		return 0, err
	}

	// Calculate fodder per cow with the fattening factor applied
	return (fodderAmount * fatteningFactor) / float64(cows), nil
}

// ValidateInputAndDivideFood validates the number of cows before dividing the food.
func ValidateInputAndDivideFood(calc FodderCalculator, cows int) (float64, error) {
	// First validate the number of cows
	if cows <= 0 {
		return 0, errors.New("invalid number of cows")
	}

	// If validation passes, call DivideFood
	return DivideFood(calc, cows)
}

// SillyNephewError represents errors caused by the silly nephew.
type SillyNephewError struct {
	cows int
}

// Error returns the error message for SillyNephewError.
func (e *SillyNephewError) Error() string {
	if e.cows < 0 {
		return fmt.Sprintf("%d cows are invalid: there are no negative cows", e.cows)
	}
	return fmt.Sprintf("%d cows are invalid: no cows don't need food", e.cows)
}

// ValidateNumberOfCows validates the number of cows and returns a custom error if invalid.
func ValidateNumberOfCows(cows int) error {
	if cows <= 0 {
		return &SillyNephewError{cows}
	}
	return nil
}
