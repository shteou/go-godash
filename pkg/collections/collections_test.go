package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	result := Map([]int{1, 2, 3}, func(x int) int {
		return x + 10
	})

	assert.Equal(t, []int{11, 12, 13}, result)
}

func TestReduce(t *testing.T) {
	// Sum reduce
	result := Reduce([]int{1, 2, 3}, 0, func(x int, a int) int {
		return a + x
	})

	assert.Equal(t, 6, result)
}

func TestSum(t *testing.T) {
	result := Sum([]int{1, 2, 3})

	assert.Equal(t, 6, result)
}

func TestFilter(t *testing.T) {
	result := Filter([]int{1, 2, 3}, func(x int) bool {
		return x > 1
	})

	assert.Equal(t, []int{2, 3}, result)
}

func TestTake(t *testing.T) {
	result := Take([]int{1, 2, 3}, 2)

	assert.Equal(t, []int{1, 2}, result)
}

func TestTakeWhile(t *testing.T) {
	result := TakeWhile([]int{1, 2, 3}, func(x int) bool {
		return x < 2
	})

	assert.Equal(t, []int{1}, result)
}

func TestTakeWhileAll(t *testing.T) {
	result := TakeWhile([]int{1, 2, 3}, func(x int) bool {
		return true
	})

	assert.Equal(t, []int{1, 2, 3}, result)
}

func TestTakeWhileNone(t *testing.T) {
	result := TakeWhile([]int{1, 2, 3}, func(x int) bool {
		return false
	})

	assert.Equal(t, []int{}, result)
}

func TestDropIdentity(t *testing.T) {
	result := Drop([]int{1, 2, 3}, 0)

	assert.Equal(t, []int{1, 2, 3}, result)
}

func TestDrop(t *testing.T) {
	result := Drop([]int{1, 2, 3}, 1)

	assert.Equal(t, []int{2, 3}, result)
}

func TestDropWhileIdentity(t *testing.T) {
	result := DropWhile([]int{1, 2, 3}, func(x int) bool {
		return false
	})

	assert.Equal(t, []int{1, 2, 3}, result)
}

func TestDropWhile(t *testing.T) {
	result := DropWhile([]int{1, 2, 3}, func(x int) bool {
		return x < 2
	})

	assert.Equal(t, []int{2, 3}, result)
}
