// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"fmt"
	"math"
	"strconv"
	"strings"
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

// StringHandler receives a string and calls Add, Subtract, Multiply or Divide based on its content.
func StringHandler(input string) (ret float64, err error) {
	operator := strings.IndexAny(input, "+-*/")
	if operator == -1 {
		return 0, fmt.Errorf("unrecognized operation")
	}
	elements := strings.Split(input, string(input[operator]))
	if len(elements) != 2 {
		return 0, fmt.Errorf("not two parts")
	}
	a, err := strconv.ParseFloat(strings.TrimSpace(elements[0]), 64)
	if err != nil {
		return 0, fmt.Errorf("a: failed parsing float64")
	}
	b, err := strconv.ParseFloat(strings.TrimSpace(elements[1]), 64)
	if err != nil {
		return 0, fmt.Errorf("b: failed parsing float64")
	}
	switch string(input[operator]) {
	case "+":
		return Add(a, b), nil
	case "-":
		return Subtract(a, b), nil
	case "*":
		return Multiply(a, b), nil
	case "/":
		ret, err = Divide(a, b)
		return ret, err
	}
	return 0, fmt.Errorf("failed to decide which function to call")
}
