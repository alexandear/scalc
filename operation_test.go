package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPerformOperationIn(t *testing.T) {
	testPerformOperation(t, PerformOperationInef)
}

func TestPerformOperation(t *testing.T) {
	testPerformOperation(t, PerformOperation)
}

func testPerformOperation(t *testing.T, performFunc func(opFunc OpFunc, n uint, sets []Iterator) Iterator) {
	for name, tc := range map[string]struct {
		opFunc   OpFunc
		n        uint
		sets     []Iterator
		expected Iterator
	}{
		"LE N=1 one set": {
			opFunc:   OpFuncLE,
			n:        1,
			sets:     []Iterator{NewIterableSlice([]int{1, 2, 3})},
			expected: NewIterableSlice([]int{}),
		},
		"LE N=1 two sets": {
			opFunc: OpFuncLE,
			n:      1,
			sets: []Iterator{
				NewIterableSlice([]int{1, 2, 3}),
				NewIterableSlice([]int{2, 3, 4}),
			},
			expected: NewIterableSlice([]int{}),
		},
		"LE N>1 one set": {
			opFunc:   OpFuncLE,
			n:        3,
			sets:     []Iterator{NewIterableSlice([]int{1, 2, 3})},
			expected: NewIterableSlice([]int{1, 2, 3}),
		},
		"LE N=2 two sets equal size": {
			opFunc: OpFuncLE,
			n:      2,
			sets: []Iterator{
				NewIterableSlice([]int{1, 2, 3}),
				NewIterableSlice([]int{2, 3, 4}),
			},
			expected: NewIterableSlice([]int{1, 4}),
		},
		"LE N=2 two sets small size=1": {
			opFunc: OpFuncLE,
			n:      2,
			sets: []Iterator{
				NewIterableSlice([]int{1}),
				NewIterableSlice([]int{2}),
			},
			expected: NewIterableSlice([]int{1, 2}),
		},
		"LE N=2 two sets different size": {
			opFunc: OpFuncLE,
			n:      2,
			sets: []Iterator{
				NewIterableSlice([]int{1}),
				NewIterableSlice([]int{2, 3}),
			},
			expected: NewIterableSlice([]int{1, 2, 3}),
		},
		"LE N=2 three sets": {
			opFunc: OpFuncLE,
			n:      2,
			sets: []Iterator{
				NewIterableSlice([]int{1, 2, 4, 5, 6}),
				NewIterableSlice([]int{-1, 0, 3, 5}),
				NewIterableSlice([]int{1, 6, 7, 8, 9, 10}),
			},
			expected: NewIterableSlice([]int{-1, 0, 2, 3, 4, 7, 8, 9, 10}),
		},
		"GR N=1 two sets first has less size than a second": {
			opFunc: OpFuncGR,
			n:      1,
			sets: []Iterator{
				NewIterableSlice([]int{2, 3, 4}),
				NewIterableSlice([]int{1, 2, 3, 4, 5}),
			},
			expected: NewIterableSlice([]int{2, 3, 4}),
		},
		"GR N=1 two sets second has less size than a first": {
			opFunc: OpFuncGR,
			n:      1,
			sets: []Iterator{
				NewIterableSlice([]int{1, 2, 3, 4, 5}),
				NewIterableSlice([]int{2, 3}),
			},
			expected: NewIterableSlice([]int{2, 3}),
		},
		"GR N=1 three sets": {
			opFunc: OpFuncGR,
			n:      1,
			sets: []Iterator{
				NewIterableSlice([]int{1, 2, 4, 5, 6}),
				NewIterableSlice([]int{-1, 0, 3, 5}),
				NewIterableSlice([]int{1, 6, 7, 8, 9, 10}),
			},
			expected: NewIterableSlice([]int{1, 5, 6}),
		},
		"EQ N=1 one set": {
			opFunc:   OpFuncEQ,
			n:        1,
			sets:     []Iterator{NewIterableSlice([]int{1, 2, 3})},
			expected: NewIterableSlice([]int{1, 2, 3}),
		},
		"EQ N=3 three sets": {
			opFunc: OpFuncEQ,
			n:      3,
			sets: []Iterator{
				NewIterableSlice([]int{1, 2, 3}),
				NewIterableSlice([]int{1, 2, 3}),
				NewIterableSlice([]int{2, 3}),
			},
			expected: NewIterableSlice([]int{2, 3}),
		},
		"EQ N=1 three sets": {
			opFunc: OpFuncEQ,
			n:      1,
			sets: []Iterator{
				NewIterableSlice([]int{1, 2, 4, 5, 6}),
				NewIterableSlice([]int{-1, 0, 3, 5}),
				NewIterableSlice([]int{1, 6, 7, 8, 9, 10}),
			},
			expected: NewIterableSlice([]int{-1, 0, 2, 3, 4, 7, 8, 9, 10}),
		},
	} {
		t.Run(name, func(t *testing.T) {
			actual := performFunc(tc.opFunc, tc.n, tc.sets)

			assert.Equal(t, tc.expected, actual)
		})
	}
}
