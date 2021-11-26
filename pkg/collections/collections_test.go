package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	result := Map([]int{1, 2, 3}, func(x int) int {
		return x + 10
	})

	assert.Equal(t, result, []int{11, 12, 13})
}

func TestReduce(t *testing.T) {
	// Sum reduce
	result := Reduce([]int{1, 2, 3}, 0, func(x int, a int) int {
		return a + x
	})

	assert.Equal(t, result, 6)
}

func TestSum(t *testing.T) {
	result := Sum([]int{1, 2, 3})

	assert.Equal(t, 6, result)
}
