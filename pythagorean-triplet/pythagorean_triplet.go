package pythagorean

import (
	"math"
)

type Triplet [3]int

func Range(min, max int) []Triplet {
	result := make([]Triplet, 0)
	for i := min; i <= max-3; i++ {
		for j := i + 1; j <= max; j++ {
			mult := i*i + j*j
			sqrt := math.Sqrt((float64)(mult))
			if (int)(sqrt) <= max && (int)(sqrt)*(int)(sqrt) == mult {
				result = append(result, Triplet{i, j, int(sqrt)})
			}
		}
	}
	return result
}

func Sum(p int) []Triplet {
	result := make([]Triplet, 0)
	trip := Range(1, p)
	for i := 0; i < len(trip); i++ {
		if trip[i][0]+trip[i][1]+trip[i][2] == p {
			result = append(result, trip[i])
		}
	}
	return result
}
