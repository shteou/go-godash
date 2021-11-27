package math

import (
	"github.com/shteou/go-godash/pkg/constraints"
	"github.com/shteou/go-godash/pkg/collections"
)

func Sum[T constraints.Numeric](xs []T) T {
	var initial T
	return collections.Reduce(xs, initial, func(x T, a T) T {
		return a + x
	})
}
