package arrays

import (
	"github.com/shteou/go-godash/pkg/types"
)

func Map[T any, U any](xs []T, f types.Mapper[T, U]) []U {
	mapped := make([]U, len(xs))

	for i, x := range xs {
		mapped[i] = f(x)
	}

	return mapped
}

func Reduce[T any, U any](xs []T, initial U, f types.Reducer[T, U]) U {
	accumulator := initial
	for _, x := range xs {
		accumulator = f(x, accumulator)
	}
	return accumulator
}


func Filter[T any](xs []T, f types.Predicate[T]) []T {
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

func TakeWhile[T any](xs []T, f types.Predicate[T]) []T {
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

func DropWhile[T any](xs []T, f types.Predicate[T]) []T {
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

func Partition[T any](xs []T, f types.Predicate[T]) ([]T, []T) {
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

func Reversed[T any](xs []T) []T {
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

func FindPredicate[T any](xs []T, f types.Predicate[T]) (*T, int) {
	for i, x := range xs {
		if f(x) {
			// Explicitly copy
			copy := x
			return &copy, i
		}
	}

	return nil, -1
}

// TODO: Rethink/Optimise all set operations
// Could support an arbitrary number of args,
// or guarantee order of elements, would need to
// stop converting to maps to achieve this
func setFromArray[T comparable](xs []T) map[T]bool {
	xsSet := map[T]bool{}
	for _, x := range xs {
	    xsSet[x] = true
	}
	return xsSet
}

func Intersection[T comparable](xs []T, ys []T) []T {
	xsSet := setFromArray(xs)
	ysSet := setFromArray(ys)

	result := []T{}

	for k, _ := range xsSet {
		if _, ok := ysSet[k]; ok {
			result = append(result, k)
		}
	}

	return result
}

func Difference[T comparable](xs []T, ys []T) []T {
	xsSet := setFromArray(xs)
	ysSet := setFromArray(ys)

	result := []T{}

	for k, _ := range xsSet {
		if _, ok := ysSet[k]; ok {
			continue
		} else {
			result = append(result, k)
		}
	}

	for k, _ := range ysSet {
		if _, ok := xsSet[k]; ok {
			continue
		} else {
			result = append(result, k)
		}
	}

	return result
}

func Union[T comparable](xs []T, ys []T) []T {
	xsSet := setFromArray(xs)
	ysSet := setFromArray(ys)

	result := xs

	for k, _ := range ysSet {
		if _, ok := xsSet[k]; ok {
			continue
		} else {
			result = append(result, k)
		}
	}

	return result
}

func Chunk[T any](xs []T, chunkSize int) [][]T {
	result := [][]T{}

	// Full chunks
	for i := 0; i < len(xs) / chunkSize; i++ {
		newChunk := []T{}
		result = append(result, []T{})
		for j := 0; j < chunkSize; j++ {
			newChunk = append(newChunk, xs[i*chunkSize+j])
		}
		result[i] = append(result[i], newChunk...)
	}

	// Remaining chunk
	remaining := len(xs) % chunkSize
	if len(xs) % chunkSize > 0 {
		newChunk := []T{}

		for i := 0; i < remaining; i++ {
			newChunk = append(newChunk, xs[len(xs)-remaining+i])
		}
		result = append(result, newChunk)
	}

	return result
}

func Flatten[T any](xs [][]T) []T {
	result := []T{}

	for i := 0; i < len(xs); i++ {
		for j := 0; j < len(xs[i]); j++ {
			result = append(result, xs[i][j])
		}
	}

	return result
}

func Without[T any](xs []T, f types.Predicate[T]) []T {
	taken := []T{}
	for _, x := range xs {
		if !f(x) {
			taken = append(taken, x)
		}
	}
	return taken
}

func Every[T any](xs []T, f types.Predicate[T]) bool {
	for _, x := range xs {
		if !f(x) {
			return false
		}
	}
	return true
}

func All[T any](xs []T, f types.Predicate[T]) bool {
	return Every(xs, f)
}

func Any[T any](xs []T, f types.Predicate[T]) bool {
	for _, x := range xs {
		if f(x) {
			return true
		}
	}
	return false
}

func Some[T any](xs []T, f types.Predicate[T]) bool {
	return Any(xs, f)
}

func FlatMap[T any, U any](xs []T, f func(T) []U) []U {
	// TODO: Optimise. It should be possible to flatten each value as
	// we iterate, rather than collecting everything first. We can save
	// some allocations and improve locality
	mapped := Map(xs, f)
	return Flatten(mapped)
}
