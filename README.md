# go-godash

An experimental generic functional utility library inspired by Lodash.

This package aims to implement many of the utility methods from Lodash
over Go arrays and Maps.

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
* Reverse
* Find
* FindPredicate

Numeric functions

* Sum
* Mean


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
