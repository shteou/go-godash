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
