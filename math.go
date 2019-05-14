// Package badmath implements math functions for accounting purposes
package badmath

// Add returns the sum of two or more ints
func Add(nums ...int) int {
	total := 0
	for _, x := range nums {
		total += x
	}

	return total
}

// Multiply returns the product of x times y
func Multiply(x, y int) int {
	total := 0
	for i := 0; i < x; i++ {
		total = Add(total, y)
	}

	return total
}
