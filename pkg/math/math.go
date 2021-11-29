package math

import (
	"github.com/shteou/go-godash/pkg/arrays"
	"github.com/shteou/go-godash/pkg/constraints"
)

func Sum[T constraints.Numeric](xs []T) T {
	return arrays.Reduce(xs, T(0), func(x T, a T) T {
		return a + x
	})
}

func Mean[T constraints.Numeric](xs []T) T {
	return Sum(xs) / T(len(xs))
}

func Average[T constraints.Numeric](xs []T) T {
	return Mean(xs)
}

func Mode[T comparable](xs []T) T {
	counts := map[T]int{}

	for i, x := range xs {
		if val, ok := counts[x]; ok {
			counts[x] = counts[x] + 1
		} else {
			counts[x] = 1
		}
	}

	var mostCommonCount := 0
	var mostCommonKey T
	for k, v := range counts {
		if v > mostCommonCount {
			mostCommonCount = v
			mostCommonKey = k
		}
	}

	return mostCommonKey
}