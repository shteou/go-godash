package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	result := Sum([]int{1, 2, 3})

	assert.Equal(t, 6, result)
}

func TestAverage(t *testing.T) {
	result := Average([]int{2, 2, 5})

	assert.Equal(t, 3, result)
}

func TestMode(t *testing.T) {
	result := Mode([]int{2, 3, 3, 4})

	assert.Equal(t, 3, result)
}

func TestModeString(t *testing.T) {
	result := Mode([]string{"a", "b", "c", "b", "c"})

	assert.Contains(t, []string{"b", "c"}, result)
}
