package collections

import "github.com/shteou/go-godash/pkg/constraints"

func Map[T any, U any](xs []T, f func(T) U) []U {
	mapped := make([]U, len(xs))

	for i, x := range xs {
		mapped[i] = f(x)
	}

	return mapped
}

func Reduce[T any, U any](xs []T, initial U, f func(T, U) U) U {
	accumulator := initial
	for _, x := range xs {
		accumulator = f(x, accumulator)
	}
	return accumulator
}

func Sum[T constraints.Numeric](xs []T) T {
	return Reduce(xs, T(0), func(x T, a T) T {
		return a + x
	})
}