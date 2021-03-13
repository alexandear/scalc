package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPerformOperation(t *testing.T) {
	for name, tc := range map[string]struct {
		opFunc   func(idx, n uint) bool
		n        uint
		sets     [][]int
		expected []int
	}{
		"LE 2 a": {
			opFunc:   OpFuncLE,
			n:        2,
			sets:     [][]int{{1, 2, 3}, {2, 3, 4}},
			expected: []int{1, 4},
		},
		"GR 1 b c": {
			opFunc:   OpFuncGR,
			n:        1,
			sets:     [][]int{{2, 3, 4}, {1, 2, 3, 4, 5}},
			expected: []int{2, 3, 4},
		},
		"GR 1 c": {
			opFunc:   OpFuncGR,
			n:        1,
			sets:     [][]int{{1, 2, 3, 4, 5}, {2, 3}},
			expected: []int{2, 3},
		},
		"EQ 3 a a b": {
			opFunc:   OpFuncEQ,
			n:        3,
			sets:     [][]int{{1, 2, 3}, {1, 2, 3}, {2, 3, 4}},
			expected: []int{2, 3},
		},
	} {
		t.Run(name, func(t *testing.T) {
			actual := PerformOperation(tc.opFunc, tc.n, tc.sets...)

			assert.Equal(t, tc.expected, actual)
		})
	}
}
