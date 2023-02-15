package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.
type SillyNephewError struct {
	cows int
}

func (e *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", e.cows)
}

// TODO: Define the SillyNephewError type here.

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {

	value, err := weightFodder.FodderAmount()
	var ErrNegativeValue = errors.New("negative fodder")
	var ErrDivisionByZero = errors.New("division by zero")
	var NewPhewError = SillyNephewError{cows: cows}
	if cows == 0 {
		return 0, ErrDivisionByZero
	}
	if err != nil {
		if err == ErrScaleMalfunction {
			if value < 0 {
				return 0, ErrNegativeValue
			}
			if cows < 0 {
				return 0, &NewPhewError
			}
			return (value * 2.0) / float64(cows), nil
		} else {
			return 0, err
		}
	}
	if value < 0 {
		return 0, ErrNegativeValue
	}
	if cows < 0 {
		return 0, &NewPhewError
	}
	return value / float64(cows), nil
}
