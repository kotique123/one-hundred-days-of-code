// Package triangle is a small library that contains a function to determine a kind of a triangle.
package triangle

// Kind is a datatype to represent a type of a triangle.
type Kind string

// These constants represent types of a triangle
const (
	NaT = "NaT" // not a triangle
	Equ = "Equ" // equilateral
	Iso = "Iso" // isosceles
	Sca = "Sca" // scalene
)

// KindFromSides determines a Kind of a triangle based on length of three sides
func KindFromSides(a, b, c float64) Kind {
	if a <= 0 || b <= 0 || c <= 0 {
		return NaT
	}
	if a+b >= c && b+c >= a && a+c >= b {
		if a == b && b == c {
			return Equ
		} else if a == b || a == c || b == c {
			return Iso
		} else {
			return Sca
		}
	}
	return NaT
}
