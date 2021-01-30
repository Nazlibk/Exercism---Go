package triangle

import "math"

type Kind string

const (
	NaT = "NaT" // not a triangle
	Equ = "Equ" // equilateral
	Iso = "Iso" // isosceles
	Sca = "Sca" // scalene
	Deg = "Deg" //degenerate
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	if a <= 0 || b <= 0 || c <= 0 || a+b < c || b+c < a || a+c < b ||
		math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) ||
		math.IsInf(a, 1) || math.IsInf(b, 1) || math.IsInf(c, 1) ||
		math.IsInf(a, -1) || math.IsInf(b, -1) || math.IsInf(c, -1) {
		return NaT
	}
	if a == b && b == c && a == c {
		return Equ
	}
	if a+b == c || b+c == a || a+c == b {
		return Deg
	}
	if a == b || b == c || a == c {
		return Iso
	}

	return Sca
}
