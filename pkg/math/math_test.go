package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	result := Sum([]int{1, 2, 3})

	assert.Equal(t, 6, result)
}

func TestSumComplex(t *testing.T) {
	var result complex64
	result = Sum([]complex64{
		complex(float32(34.0), float32(15.5)),
		complex(float32(12), float32(-3))})

	assert.Equal(t, float32(46), real(result))
	assert.Equal(t, float32(12.5), imag(result))
}
