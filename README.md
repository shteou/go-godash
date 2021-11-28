# go-godash

An experimental generic functional library for transforming collections
of data, inspired by packages such as Lodash.

This package aims to implement many of the methods from Lodash over
Go arrays and Maps. All functions in this module are intended to work with
native Go types (i.e. do not define their own containers/iterators), and
should be type safe at compile-time.

The functions defined here are only loosely inspired by Lodash. Sometimes
they deviate because of impedance between the two languages, or sometimes
because I didn't bother to read the documentation properly.

_Caveat_: I have very little experience with Go generics, so the API will
evolve heavily. It may never hit version 1, depending on how collections
evolve in Go once generics lands.

## Example

Here's a contrived example of implementing and `incrementAndSum` function.
The examples themselves are defined for a concrete type, but the underlying
`Sum` and `Map` functions are generic across `Numeric` and `all` types respectively.

```go
// A non-generic increment function for integers
func increment(x int) int {
	return x + 1
}

// A non-generic increment and sum function for arrays of integers
func incrementAndSum(arr []int) int {
	// Utilises generic Map and Sum functions
	result := Sum(Map(arr, increment))
}

// Returns 9
incrementAndSum([]int{1, 2, 3})
```

## Implemented functions

Arrays:

* Map
* Reduce
* Filter
* Take
* TakeWhile
* Drop
* DropWhile
* Partition
* First (Head)
* Last
* Reversed
* Find
* FindPredicate
* Intersection
* Difference
* Chunk
* Flatten
* Without

Numeric functions:

* Sum
* Mean

### Not implemented

Some functions are unimplemented because they do not translate in a type-safe manner.

* FlattenDeep/FlattenDepth - It's not immediately obvious to me how to achieve this
  in a type safe way, so these functions are omitted
* FromPairs - Assumes an input array of untyped data which cannot be achieved at
  compile time

## Limitations

Complex types are not supported on operations for Numeric types (Sum, Mean etc.).  
I was able to find solutions to zero-initialise complex types for Sum, but this
technique does not translate to Mean (and probably other generic Numeric operations) where
we need to divide an arbitrary complex type (complex64 or 128) by the length of
the array. There is no way to infer whether it should be divided by a float32 or float64
type.

See [this section](https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#no-way-to-express-convertibility) and [this section](https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#no-association-between-float-and-complex) of the generics proposal for more information.

## gotip

This library uses generics which is not yet available in go. You'll need to
use [gotip](https://pkg.go.dev/golang.org/dl/gotip) to test this out.
