// Package complexnumbers is a small library that provides functionality to
// work with complex numbers in Golang.
package complexnumbers

import (
	"math"
)

// Number represents a complex number containing real and imaginary parts.
type Number struct {
	real      float64 // real part
	imaginary float64 // imaginary part
}

// Real returns a real portion of a Number.
func (n Number) Real() float64 {
	return n.real
}

// Imaginary returns an imaginary portion of a Number.
func (n Number) Imaginary() float64 {
	return n.imaginary
}

// Add combines real and imaginary parts of two complex numbers and returns a
// Number.
func (n1 Number) Add(n2 Number) Number {
	var number Number
	number.real = n1.real + n2.real
	number.imaginary = n1.imaginary + n2.imaginary
	return number

}

// Subtract subtracts real and imaginary parts of two complex numbers and
// returns a Number.
func (n1 Number) Subtract(n2 Number) Number {
	var number Number
	number.real = n1.real - n2.real
	number.imaginary = n1.imaginary - n2.imaginary
	return number
}

// Multiply multiplies two numbers according to a formula and returns a
// Number.
func (n1 Number) Multiply(n2 Number) Number {
	var number Number
	number.real = n1.real*n2.real - n1.imaginary*n2.imaginary
	number.imaginary = n1.real*n2.imaginary + n1.imaginary*n2.real
	return number
}

// Times multiplies two parts of a number by a factor and returns a Number.
func (n Number) Times(factor float64) Number {
	var number Number
	number.real = n.real * factor
	number.imaginary = n.imaginary * factor
	return number
}

// Divide divides two Numbers according to a formula and returns a Number.
func (n1 Number) Divide(n2 Number) Number {
	var number Number
	number.real = (n1.real*n2.real + n1.imaginary*n2.imaginary) / (n2.real*n2.real + n2.imaginary*n2.imaginary)
	number.imaginary = (n1.imaginary*n2.real - n1.real*n2.imaginary) / (n2.real*n2.real + n2.imaginary*n2.imaginary)
	return number
}

// Conjugate returns a conjugate of a Number.
func (n Number) Conjugate() Number {
	var number Number
	number.real = n.real
	number.imaginary = (-1) * n.imaginary
	return number
}

// Abs returns an abosulte value of a Number.
func (n Number) Abs() float64 {
	absolute_value := math.Sqrt(n.real*n.real + n.imaginary*n.imaginary)
	return absolute_value
}

// Exp returns an exponent of a Number. More information available at README_Exercism.md.
func (n Number) Exp() Number {
	var number Number
	number.real = math.Cos(n.imaginary)
	number.imaginary = math.Sin(n.imaginary)
	return number.Times(math.Exp(n.real))
}

// AllOperationsTogether is a function used to benchmark the solution and it
// wasn't included in original Exercism problem.
func (n1 Number) AllOperationsTogether(factor float64, n2 Number) {
	n1.Real()
	n1.Imaginary()
	n1.Add(n2)
	n1.Subtract(n2)
	n1.Multiply(n2)
	n1.Times(factor)
	n1.Divide(n2)
	n1.Conjugate()
	n1.Abs()
	n1.Exp()
}
