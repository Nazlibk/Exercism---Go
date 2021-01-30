package darts

import "math"

func Score(x, y float64) int {
	radius := math.Pow(math.Pow(x, 2)+math.Pow(y, 2), 0.5)
	if radius > 10 {
		return 0
	}
	if radius <= 10 && radius > 5 {
		return 1
	}
	if radius <= 5 && radius > 1 {
		return 5
	}
	return 10
}
