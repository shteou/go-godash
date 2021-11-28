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