package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	result := Sum([]int{1, 2, 3})

	assert.Equal(t, 6, result)
}

func TestMean(t *testing.T) {
	result := Mean([]int{2, 2, 5})

	assert.Equal(t, 3, result)
}