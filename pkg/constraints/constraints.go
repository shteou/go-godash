package constraints

type Numeric interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64 | ~complex64 | ~complex128
}
