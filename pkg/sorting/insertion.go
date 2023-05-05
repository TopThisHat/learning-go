package sorting

import (
	"golang.org/x/exp/constraints"
)

func insertion_sort[T constraints.Ordered](s []T) {
	for i := 1; i < len(s); i++ {
		current := i
		for current > 0 && s[current] < s[current-1] {
			s[current], s[current-1] = s[current-1], s[current]
			current--
		}
	}
	return
}
