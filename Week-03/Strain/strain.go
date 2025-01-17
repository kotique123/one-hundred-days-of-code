// Package strain provides utilities for filtering slices based on a given condition.
package strain

// Keep filters elements from the input slice based on the provided condition function.
// It returns a new slice containing only the elements for which the condition function returns true.
//
// S is a type parameter constrained to slices of E.
// E is a type parameter representing the element type of the slice.
//
// Parameters:
// - input: The input slice of type S to be filtered.
// - condition: A function that takes an element of type E and returns a boolean.
//
// Returns:
// A new slice of type S containing only the elements that satisfy the condition.
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	isEven := func(n int) bool { return n%2 == 0 }
//	result := Keep(numbers, isEven) // result: []int{2, 4}
func Keep[S ~[]E, E any](input S, condition func(E) bool) S {
	// helped a lot https://exercism.org/tracks/go/exercises/strain/solutions/ogv
	// thanks to him
	result := make([]E, 0, len(input))
	for _, val := range input {
		if condition(val) {
			result = append(result, val)
		}
	}
	return result
}

// Discard filters elements from the input slice based on the provided condition function.
// It returns a new slice containing only the elements for which the condition function returns false.
//
// S is a type parameter constrained to slices of E.
// E is a type parameter representing the element type of the slice.
//
// Parameters:
// - input: The input slice of type S to be filtered.
// - condition: A function that takes an element of type E and returns a boolean.
//
// Returns:
// A new slice of type S containing only the elements that do not satisfy the condition.
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	isEven := func(n int) bool { return n%2 == 0 }
//	result := Discard(numbers, isEven) // result: []int{1, 3, 5}
func Discard[S ~[]E, E any](input S, condition func(E) bool) S {
	result := make([]E, 0, len(input))
	for _, val := range input {
		if !condition(val) {
			result = append(result, val)
		}
	}
	return result
}
