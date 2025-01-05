// uncomment this file to work with it
/*
package complexnumbers

import (
	"math"
)

type Number struct {
	real      float64 // real part
	imaginary float64 // imaginary part
}

func (n Number) Real() float64 {
	return n.real
}

func (n Number) Imaginary() float64 {
	return n.imaginary
}

func (n1 Number) Add(n2 Number) Number {
	n1.real += n2.real
	n1.imaginary += n2.imaginary
	return n1
}

func (n1 Number) Subtract(n2 Number) Number {
	n1.real -= n2.real
	n1.imaginary -= n2.imaginary
	return n1
}

func (n1 Number) Multiply(n2 Number) Number {
	var number Number
	number.real = n1.real*n2.real - n1.imaginary*n2.imaginary
	number.imaginary = n1.real*n2.imaginary + n1.imaginary*n2.real
	return number
}

func (n Number) Times(factor float64) Number {
	n.real = n.real * factor
	n.imaginary = n.imaginary * factor
	return n
}

func (n1 Number) Divide(n2 Number) Number {
	var number Number
	number.real = (n1.real*n2.real + n1.imaginary*n2.imaginary) / (n2.real*n2.real + n2.imaginary*n2.imaginary)
	number.imaginary = (n1.imaginary*n2.real - n1.real*n2.imaginary) / (n2.real*n2.real + n2.imaginary*n2.imaginary)
	return number
}

func (n Number) Conjugate() Number {
	n.imaginary = (-1) * n.imaginary
	return n
}

func (n Number) Abs() float64 {
	absolute_value := math.Sqrt(n.real*n.real + n.imaginary*n.imaginary)
	return absolute_value
}

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
*/