// Package lsproduct is a small library that has allow to determine the largest
// product in a series of numbers.
package lsproduct

import (
	"errors"
	"strconv"
)

// LargestSeriesProduct accepts a series of numbers as a string and a span of
// series and returns sum and error.
func LargestSeriesProduct(digits string, span int) (int64, error) {
	// checks if span is negative
	if span < 0 {
		return 0, errors.New("span must not be negative")
		// chekcs if span is longer than series
	} else if span > len(digits) {
		return 0, errors.New("span must be smaller than string length")
	}
	// allocates memory for a slice for numbers from a string
	nums := make([]int, len(digits))
	for i := 0; i < len(digits); i++ {
		num, err := strconv.Atoi(string(digits[i]))
		if err == nil {
			nums[i] = num
			// if a sequence is not valid, meaning that there are values other than
			// numerical, returns an error
		} else {
			return 0, errors.New("digits input must only contain digits")
		}
	}
	start := 0
	end := span
	var sum int
	var max int
	for start < len(nums)-1 {
		temp_slice := nums[start:end]
		sum = 1
		for i := 0; i < len(temp_slice); i++ {
			sum *= (temp_slice[i])
		}
		start++
		end = span + start
		if sum > max {
			max = sum
		}
		if end > len(nums) {
			break
		}
	}
	return int64(max), nil
}
