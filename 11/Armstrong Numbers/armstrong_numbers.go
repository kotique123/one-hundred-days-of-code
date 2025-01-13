// Package armstrong is a small library that has functionality to determine if
// a number is an armstrong number.

package armstrong

import (
	"math"
)

// GetNumberOfDigitsRecursive returns number of digits in a number using
// recursive approach.
func GetNumberOfDigitsRecursive(number int) int {
	if number < 10 {
		return 1
	}
	return 1 + GetNumberOfDigitsRecursive(number/10)
}

// GetNumberOfDigitsRecursive returns number of digits in a number using
// loop-based appraoch.
func GetNumberOfDigitsLoop(number int) int {
	counter := 0
	for number > 0 {
		number = number / 10
		counter++
	}
	return counter
}

// SplitNumberByDigits returns a slice of digits based on an input number and a
// number of digits.
func SplitNumberByDigits(number, num_of_digits int) []int {
	digits := make([]int, 0, num_of_digits)
	for i := 1; i <= num_of_digits; i++ {
		power := int(math.Pow(10, float64(num_of_digits-i)))
		digits = append(digits, number/power)
		number = number % power
	}
	return digits
}

// ArmstrongSum retuns a sum of digits of a number using Armstrong method.
// More at https://en.wikipedia.org/wiki/Narcissistic_number
func ArmstrongSum(digits []int, number_of_digits int) int {
	sum := 0.0
	for _, digit := range digits {
		sum += math.Pow(float64(digit), float64(number_of_digits))
	}
	return int(sum)
}

// IsNumber returns a boolean value. It returns true if an input number
// is an Armstrong number.
func IsNumber(n int) bool {
	num_of_digits := GetNumberOfDigitsLoop(n)
	return n == ArmstrongSum(SplitNumberByDigits(n, num_of_digits), num_of_digits)
}
