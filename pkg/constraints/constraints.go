package constraints

// Numeric defines all types which are simple numerics.
// That is integer and floating point values, and
// excluding complex values.
type Numeric interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
}
