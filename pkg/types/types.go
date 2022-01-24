package types

// Mapper is a function which accepts one type
// and returns another.
type Mapper[T any, U any] func(T) U

// Mapper is a function which accepts one type
// and returns another.
type MapperWithError[T any, U any] func(T) (U, error)

// Predicate is a function which accepts one type
// and returns a bool indicating whether the predicate
// was satisfied.
type Predicate[T any] func(T) bool

// Reducer is a function which accepts a parameter of
// type T and a parameter of type U, the accumulator
// and returns a new value of type U (this return value
// is passed with the next element in the Reduce
// function).
type Reducer[T any, U any] func(T, U) U
