package domain

import (
	"math"
)

func Assessment(nbc int, age int) []int {
	asmt := []int{}
	for i := 0; i < nbc; i++ {
		cycle := int(math.Pow(2, float64(i)))
		if age%cycle == 0 {
			asmt = append(asmt, i)
		}
	}
	return asmt
}
