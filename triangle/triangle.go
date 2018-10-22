package triangle

import "math"

// Kind holds a numerical representation of a type of triangle
type Kind int

const (
	// NaT is a constant used when the given values do not make up a triangle
	NaT Kind = iota

	// Equ is a constant representing an Equilateral triangle
	Equ

	// Iso is a constant representing an Isosceles triangle
	Iso

	// Sca is a constant representing a Scalene triangle
	Sca
)

// KindFromSides returns the kind of triangle given lengths a, b, and c.
func KindFromSides(a, b, c float64) Kind {

	if !isTriangle(a, b, c) {
		return NaT
	}

	if isEquilateral(a, b, c) {
		return Equ
	}

	if isScalene(a, b, c) {
		return Sca
	}

	return Iso
}

func isEquilateral(a, b, c float64) bool {
	return a == b && b == c
}

func isScalene(a, b, c float64) bool {
	return a != b && a != c && b != c
}

func isTriangle(a, b, c float64) bool {

	if !areValidSides(a, b, c) {
		return false
	}

	if !arePositiveSides(a, b, c) {
		return false
	}

	if !triangleInequalityHolds(a, b, c) {
		return false
	}

	return true
}

func areValidSides(a, b, c float64) bool {
	sides := []float64{a, b, c}
	const EitherInfinity = 0

	for _, val := range sides {
		if math.IsNaN(val) || math.IsInf(val, EitherInfinity) {
			return false
		}
	}

	return true
}

func arePositiveSides(a, b, c float64) bool {
	return a > 0 && b > 0 && c > 0
}

func triangleInequalityHolds(a, b, c float64) bool {
	return (a+b >= c) && (a+c >= b) && (b+c >= a)
}
