package types

type Mapper[T any, U any] func(T) U
type Predicate[T any] func(T) bool
type Reducer[T any, U any] func(T, U) U