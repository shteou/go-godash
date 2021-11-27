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

func Filter[T any](xs []T, f func(T) bool) []T {
	taken := []T{}
	for _, x := range xs {
		if f(x) {
			taken = append(taken, x)
		}
	}
	return taken
}

func Take[T any](xs []T, n int) []T {
	taken := make([]T, n)

	for i := 0; i<n; i++ {
		taken[i] = xs[i]
	}

	return taken
}

func TakeWhile[T any](xs []T, f func(T) bool) []T {
	taken := []T{}

	for _, x := range xs {
		if f(x) {
			taken = append(taken, x)
		} else {
			break
		}
	}

	return taken
}

func Drop[T any](xs []T, n int) []T {
	taken := make([]T, len(xs)-n)

	for i := n; i<len(xs); i++ {
		taken[i-n] = xs[i]
	}

	return taken
}

func DropWhile[T any](xs []T, f func(T) bool) []T {
	taken := []T{}
	dropping := true

	for _, x := range xs {
		if !f(x) {
			dropping = false
		}

		if dropping {
			continue
		} else {
			taken = append(taken, x)
		}
	}

	return taken
}

func Partition[T any](xs []T, f func(T) bool) ([]T, []T) {
	passed := []T{}
	failed := []T{}

	for _, x := range xs {
		if f(x) {
			passed = append(passed, x)
		} else {
			failed = append(failed, x)
		}
	}

	return passed, failed
}

func Head[T any](xs []T) T {
	return xs[0]
}

func First[T any](xs []T) T {
	return Head(xs)
}

func Last[T any](xs []T) T {
	return xs[len(xs)-1]
}

func Reverse[T any](xs []T) []T {
	reversed := make([]T, len(xs))

	for i, _ := range xs {
		reversed[len(xs)-i-1] = xs[i]
	}

	return reversed
}

func Find[T comparable](xs []T, t T) (*T, int) {
	for i, x := range xs {
		if x == t {
			// Explicitly copy
			copy := x
			return &copy, i
		}
	}
	return nil, -1
}

func FindPredicate[T any](xs []T, f func(T) bool) (*T, int) {
	for i, x := range xs {
		if f(x) {
			// Explicitly copy
			copy := x
			return &copy, i
		}
	}

	return nil, -1
}