package domain

import (
	"math"
)

func Assessment(nbc int, day int) []int {
	asmt := []int{}
	for i := 1; i <= nbc; i++ {
		cycle := int(math.Pow(float64(i), 2))
		mod := day % cycle
		if mod == 1 {
			asmt = append(asmt, i)
		}
	}
	return asmt
}

func NeedCleanLastCell(nbc int, day int) bool {
	var cycle int
	cycle = 2 ^ nbc
	mod := day % cycle
	return mod == 0
}
