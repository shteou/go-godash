package math

import (
	"github.com/shteou/go-godash/pkg/arrays"
	"github.com/shteou/go-godash/pkg/constraints"
)

// Sum adds all the elements in the supplied array xs
// together, returning the result.
func Sum[T constraints.Numeric](xs []T) T {
	return arrays.Reduce(xs, T(0), func(x T, a T) T {
		return a + x
	})
}

// Mean returns the average value of all elements in the
// supplied array xs.
func Mean[T constraints.Numeric](xs []T) T {
	return Sum(xs) / T(len(xs))
}

// Average returns the average value of all elements in the
// supplied array xs.
// Average is an alias of Mean.
func Average[T constraints.Numeric](xs []T) T {
	return Mean(xs)
}

// Mode returns the most common element in the supplied array
// xs. If multiple elements are as common as eachother, any
// of those values can be returned.
func Mode[T comparable](xs []T) T {
	counts := map[T]int{}

	for _, x := range xs {
		if val, ok := counts[x]; ok {
			counts[x] = val + 1
		} else {
			counts[x] = 1
		}
	}

	var mostCommonCount int
	var mostCommonKey T
	for k, v := range counts {
		if v > mostCommonCount {
			mostCommonCount = v
			mostCommonKey = k
		}
	}

	return mostCommonKey
}