// Package badmath implements math functions for accounting purposes
package badmath

import "fmt"

// Add returns the sum of two or more ints
func Add(nums ...int) (int, error) {
	if len(nums) < 2 {
		return 0, fmt.Errorf("Not enough arguments")
	}
	total := 0
	for _, x := range nums {
		total += x
	}

	log("Add()", total, nums...)

	return total, nil
}

// Multiply returns the product of x times y
func Multiply(x, y int) int {
	total := 0
	for i := 0; i < x; i++ {
		total, _ = Add(total, y)
	}

	log("Multiply()", total, x, y)

	return total
}

// Divide returns the result of x divided by y
func Divide(x, y int) (float64, error) {
	if y == 0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}

	total := float64(x) / float64(y)
	log("Divide()", total, x, y)
	return total, nil
}
