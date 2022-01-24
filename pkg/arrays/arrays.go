package arrays

import (
	"github.com/shteou/go-godash/pkg/types"
)

// Map creates a new array of of values by passing each
// element of xs to the Mapper function f.
func Map[T any, U any](xs []T, f types.Mapper[T, U]) []U {
	mapped := make([]U, len(xs))

	for i, x := range xs {
		mapped[i] = f(x)
	}

	return mapped
}

// Map creates a new array of values by passing each element
// of xs to the MapperWithError function f. This function short
// circuits if f returns an error, returning the elements so far
// NOTE: The output array may be of the same size as the input
// array, xs, but may not necessarily be fully populated
func MapWithError[T any, U any](xs []T, f types.MapperWithError[T, U]) ([]U, error) {
	mapped := make([]U, len(xs))

	for i, x := range xs {
		res, err := f(x)
		if err != nil {
			return mapped, err
		} else {
			mapped[i] = res
		}
	}

	return mapped, nil
}

// Reduce produces a value from xs by accumulating
// the result of each element as passed through the Reducer
// function f. The first element is passed to the Reducer with
// the supplied initial value.
func Reduce[T any, U any](xs []T, initial U, f types.Reducer[T, U]) U {
	accumulator := initial
	for _, x := range xs {
		accumulator = f(x, accumulator)
	}
	return accumulator
}

// Filter returns a new array of all elements in xs that
// pass the supplied Predicate f. If the predicate returns
// true the value is included in the resulting array, otherwise
// it is omitted.
func Filter[T any](xs []T, f types.Predicate[T]) []T {
	taken := []T{}
	for _, x := range xs {
		if f(x) {
			taken = append(taken, x)
		}
	}
	return taken
}

// Take returns the first n elements of the array xs.
func Take[T any](xs []T, n int) []T {
	taken := make([]T, n)

	for i := 0; i < n; i++ {
		taken[i] = xs[i]
	}

	return taken
}

// TakeWhile returns a new array containing the first elements
// of xs which pass the supplied Predicate f. Once the Predicate
// returns false that element and all subsequent elements are
// discarded.
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

// Drop returns a new array with the first n elements discarded
func Drop[T any](xs []T, n int) []T {
	taken := make([]T, len(xs)-n)

	for i := n; i < len(xs); i++ {
		taken[i-n] = xs[i]
	}

	return taken
}

// DropWhile returns a new array, discarding all elements of xs
// which pass the supplied Predicate f. Once the Predicate returns
// false that element and all subsequent elements are returned.
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

// Partition returns two new arrays, the first containing all
// elements of xs that pass the supplied Predicate f, and the
// second returning the remainder.
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

// Head returns the first element of the array
func Head[T any](xs []T) T {
	return xs[0]
}

// First returns the first element of the array.
// First is an alias of Head.
func First[T any](xs []T) T {
	return Head(xs)
}

// Last returns the last element of the array.
func Last[T any](xs []T) T {
	return xs[len(xs)-1]
}

// Reversed returns a new array with the elements
// in reverse order.
func Reversed[T any](xs []T) []T {
	reversed := make([]T, len(xs))

	for i, _ := range xs {
		reversed[len(xs)-i-1] = xs[i]
	}

	return reversed
}

// Find locates the first instance of the supplied element t
// in the array xs. It returns a pointer to a copy of the value
// and its index in the array.
// If the element is not found, a nil pointer is returned along
// with an index position of -1
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

// Find locates the first instance of the supplied array xs which
// passes the supplied Predicate f. It returns a pointer to a copy
// of the value and its index in the array.
// If a passing element is not found, a nil pointer is returned along
// with an index position of -1
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

// setFromArray returns a set of the supplied array's values,
// represented as a map of keys (the set elements) and bool
// values (which have no semantic meaning)
func setFromArray[T comparable](xs []T) map[T]bool {
	// TODO: Rethink/Optimise all set operations
	// Could support an arbitrary number of args,
	// or guarantee order of elements, would need to
	// stop converting to maps to achieve this
	xsSet := map[T]bool{}
	for _, x := range xs {
		xsSet[x] = true
	}
	return xsSet
}

// Intersection performs a set intersection on the supplied
// arrays xs and ys, returning a new array of unique values.
// The order of values in the set is not guaranteed.
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

// Difference performs a set difference on the supplied
// arrays, xs and ys, returning a new array of unique values.
// The order of values in the set is not guaranteed.
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

// Union performs a set union on the supplied arrays,
// xs and ys, returning a new array of unique values.
// The order of values in the set is not guaranteed.
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

// Chunk splits the supplied array into one or more new
// arrays of size chunkSize. The last array in the
// returned value may be of length n < chunkSize
func Chunk[T any](xs []T, chunkSize int) [][]T {
	result := [][]T{}

	// Full chunks
	for i := 0; i < len(xs)/chunkSize; i++ {
		newChunk := []T{}
		result = append(result, []T{})
		for j := 0; j < chunkSize; j++ {
			newChunk = append(newChunk, xs[i*chunkSize+j])
		}
		result[i] = append(result[i], newChunk...)
	}

	// Remaining chunk
	remaining := len(xs) % chunkSize
	if len(xs)%chunkSize > 0 {
		newChunk := []T{}

		for i := 0; i < remaining; i++ {
			newChunk = append(newChunk, xs[len(xs)-remaining+i])
		}
		result = append(result, newChunk)
	}

	return result
}

// Flatten merges an array of arrays into a new array
// with one fewer levels of depth
func Flatten[T any](xs [][]T) []T {
	result := []T{}

	for i := 0; i < len(xs); i++ {
		for j := 0; j < len(xs[i]); j++ {
			result = append(result, xs[i][j])
		}
	}

	return result
}

// Without returns a new array, omitting all values from
// the supplied array xs which pass the supplied Predicate f
func Without[T any](xs []T, f types.Predicate[T]) []T {
	taken := []T{}
	for _, x := range xs {
		if !f(x) {
			taken = append(taken, x)
		}
	}
	return taken
}

// Every returns true if all elements of the supplied array xs
// pass the Predicate f. An empty array yields true.
func Every[T any](xs []T, f types.Predicate[T]) bool {
	for _, x := range xs {
		if !f(x) {
			return false
		}
	}
	return true
}

// All returns true if all elements of the supplied array xs
// pass the Predicate f. An empty array yields true.
// All is an alias of Every.
func All[T any](xs []T, f types.Predicate[T]) bool {
	return Every(xs, f)
}

// Any returns true if any element in the supplied array xs
// passes the Predicate f. An empty array yields false.
func Any[T any](xs []T, f types.Predicate[T]) bool {
	for _, x := range xs {
		if f(x) {
			return true
		}
	}
	return false
}

// Some returns true if any element in the supplied array xs
// passes the Predicate f. An empty array yields false.
// Some is an alias of Any.
func Some[T any](xs []T, f types.Predicate[T]) bool {
	return Any(xs, f)
}

// FlatMap performs a Map operation over the supplied array xs,
// where the supplied function returns an array for each element.
// It then Flattens the result. It is equivalent to calling Map
// and then Flatten, however the Mapper is required to return an array.
func FlatMap[T any, U any](xs []T, f func(T) []U) []U {
	// TODO: Optimise. It should be possible to flatten each value as
	// we iterate, rather than collecting everything first. We can save
	// some allocations and improve locality
	mapped := Map(xs, f)
	return Flatten(mapped)
}
