// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"fmt"
	"math"
)

// Add adds all the other numbers to the first.
func Add(a float64, b ...float64) float64 {
	for _, v := range b {
		a += v
	}
	return a
}

// Subtract subtracts all the other numbers from the first.
func Subtract(a float64, b ...float64) float64 {
	for _, v := range b {
		a -= v
	}
	return a
}

// Multiply multiples all the given numbers together.
func Multiply(a float64, b ...float64) float64 {
	for _, v := range b {
		a *= v
	}
	return a
}

// Divide divides the first number with all the rest.
// If any other then the first equals to 0, returns an error.
func Divide(a float64, b ...float64) (float64, error) {
	for _, v := range b {
		if v == 0 {
			return 0, fmt.Errorf("divide by 0 is not allowed")
		}
		a /= v
	}
	return a, nil
}

// Sqrt takes a number and returns with the square root of it.
// If the base is negative, gives an error.
func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, fmt.Errorf("input have to be positive")
	}
	return math.Sqrt(a), nil
}
