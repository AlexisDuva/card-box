package domain

import (
	"math"
)

func Assessment(nbc int, age int) []int {
	asmt := []int{}
	for i := 1; i <= nbc; i++ {
		cycle := int(math.Pow(float64(i), 2))
		mod := age % cycle
		if mod == 1 {
			asmt = append(asmt, i-1)
		}
	}
	return asmt
}
