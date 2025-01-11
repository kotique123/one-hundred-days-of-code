// Package sieve is a small library that allows to find prime numbers using
// the Sieve of Eratosthenes algorithm.
package sieve

// Sieve returns an array of prime numbers up until a certain limit passed
// im arguments.
func Sieve(limit int) []int {
	numbers := make([]bool, limit+1)
	counter := 0
	for i := 2; i <= limit; i++ {
		numbers[i] = true
	}
	for i := 2; i*i <= limit; i++ {
		if numbers[i] {
			for j := i * i; j <= limit; j += i {
				numbers[j] = false
			}
			counter++
		}
	}
	primes := make([]int, 0, counter)
	index := 0
	for i := 2; i < len(numbers); i++ {
		if numbers[i] {
			primes = append(primes, i)
			index++
		}
	}
	return primes
}
