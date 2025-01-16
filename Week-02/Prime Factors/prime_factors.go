// Package prime is a small library that finds prime factors of a number
package prime

import "math"

// Factors returns a slice prime factors of a number
func Factors(n int64) []int64 {
	if n <= 1 {
		return []int64{}
	}

	// Estimate the size of the slice for preallocation (at most log2(n) factors for large n).
	estimatedSize := int(math.Log2(float64(n))) + 1
	factors := make([]int64, 0, estimatedSize)

	// Divide out the factor of 2 first.
	for n%2 == 0 {
		factors = append(factors, 2)
		n /= 2
	}

	// Check for odd factors from 3 up to sqrt(n).
	for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
		for n%int64(i) == 0 {
			factors = append(factors, int64(i))
			n /= int64(i)
		}
	}

	// If n is still greater than 2, it's a prime number.
	if n > 2 {
		factors = append(factors, n)
	}

	return factors
}
