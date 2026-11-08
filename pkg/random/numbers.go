package random

import (
	"math/rand"
	"slices"
)

func GenerateLotofacilNumbers15() []int {
	s := rand.Perm(25)

	s = s[:15]

	for i := range s {
		s[i]++
	}

	slices.Sort(s)

	return s
}
