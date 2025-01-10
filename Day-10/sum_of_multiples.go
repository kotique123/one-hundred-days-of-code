// Package summultiples is a small library that can find sum of multiples.
package summultiples

import "slices"

// SumMultiples returns sum of multiples a slice limited by a limit
func SumMultiples(limit int, divisors ...int) int {
	multiples := make([]int, 0, limit) // empty slice for multiples
	for _, number := range divisors {
		if number == 0 {
			continue
		}
		status := true
		for i := 1; status; i++ {
			multiple := number * i
			if multiple < limit {
				multiples = append(multiples, multiple)
			} else {
				status = false
			}
		}
	}
	slices.Sort(multiples)                // sort multiples
	multiples = slices.Compact(multiples) // remove duplicates(does not work with not sorted slices)
	sum := 0
	for _, number := range multiples {
		sum += number
	}
	return sum
}
