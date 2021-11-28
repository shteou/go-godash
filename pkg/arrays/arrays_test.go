package arrays

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

func TestDropWhile_DontStartDroppingAgain(t *testing.T) {
	result := DropWhile([]int{1, 2, 3}, func(x int) bool {
		return x == 1 || x == 3
	})

	assert.Equal(t, []int{2, 3}, result)
}

func TestPartition(t *testing.T) {
	passed, failed := Partition([]int{1, 2, 3}, func(x int) bool {
		return x < 2
	})

	assert.Equal(t, []int{1}, passed)
	assert.Equal(t, []int{2, 3}, failed)
}

func TestHead(t *testing.T) {
	result := Head([]int{1, 2, 3})

	assert.Equal(t, 1, result)
}

func TestLast(t *testing.T) {
	result := Last([]int{1, 2, 3})

	assert.Equal(t, 3, result)
}

func TestReversed(t *testing.T) {
	result := Reversed([]int{1, 2, 3})

	assert.Equal(t, []int{3, 2, 1}, result)
}

func TestFind(t *testing.T) {
	result, index := Find([]int{1, 2, 3}, 2)

	assert.Equal(t, 2, *result)
	assert.Equal(t, 1, index)
}

func TestFindAbsent(t *testing.T) {
	result, index := Find([]int{1, 2, 3}, 4)

	assert.Nil(t, result)
	assert.Equal(t, -1, index)
}

func TestFindPredicate(t *testing.T) {
	result, index := FindPredicate([]int{1, 2, 3}, func(x int) bool {
		return x == 3
	})

	assert.Equal(t, 3, *result)
	assert.Equal(t, 2, index)
}

func TestFindPredicateAbsent(t *testing.T) {
	result, index := FindPredicate([]int{1, 2, 3}, func(x int) bool {
		return x == 4
	})

	assert.Nil(t, result)
	assert.Equal(t, -1, index)
}

func TestIntersection(t *testing.T) {
	elems := Intersection([]int{1, 2, 3}, []int{2, 3, 4})

	assert.Len(t, elems, 2)
	assert.Contains(t, elems, 2)
	assert.Contains(t, elems, 3)
}

func TestDifference(t *testing.T) {
	elems := Difference([]int{1, 2, 3}, []int{2, 3, 4})

	assert.Len(t, elems, 2)
	assert.Contains(t, elems, 1)
	assert.Contains(t, elems, 4)
}

func TestUnion(t *testing.T) {
	elems := Union([]int{1, 2, 3}, []int{2, 3, 4})

	assert.Len(t, elems, 4)
	assert.Contains(t, elems, 1)
	assert.Contains(t, elems, 2)
	assert.Contains(t, elems, 3)
	assert.Contains(t, elems, 4)
}

func TestChunk(t *testing.T) {
	chunks := Chunk([]int{1, 2, 3}, 2)

	assert.Len(t, chunks, 2)
	assert.Equal(t, []int{1, 2}, chunks[0])
	assert.Equal(t, []int{3}, chunks[1])
}

func TestChunkExact(t *testing.T) {
	chunks := Chunk([]int{1, 2, 3, 4}, 2)

	assert.Len(t, chunks, 2)
	assert.Equal(t, []int{1, 2}, chunks[0])
	assert.Equal(t, []int{3, 4}, chunks[1])
}

func TestFlatten(t *testing.T) {
	flattened := Flatten([][]int{[]int{1, 2}, []int{3, 4}})

	assert.Equal(t, flattened, []int{1, 2, 3, 4})
}

func TestWithout(t *testing.T) {
	result := Without([]int{1, 2, 3}, func(x int) bool {
		return x > 2
	})

	assert.Equal(t, []int{1, 2}, result)
}
