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
		"LE N=1 one set": {
			opFunc:   OpFuncLE,
			n:        1,
			sets:     [][]int{{1, 2, 3}},
			expected: []int{},
		},
		"LE N>1 one set": {
			opFunc:   OpFuncLE,
			n:        3,
			sets:     [][]int{{1, 2, 3}},
			expected: []int{1, 2, 3},
		},
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
		"GR 1": {
			opFunc:   OpFuncGR,
			n:        1,
			sets:     [][]int{{1, 2, 4, 5, 6}, {-1, 0, 3, 5}, {1, 6, 7, 8, 9, 10}},
			expected: []int{1, 5, 6},
		},
		"EQ 1": {
			opFunc:   OpFuncEQ,
			n:        1,
			sets:     [][]int{{1, 2, 4, 5, 6}, {-1, 0, 3, 5}, {1, 6, 7, 8, 9, 10}},
			expected: []int{-1, 0, 2, 3, 4, 7, 8, 9, 10},
		},
		"LE 2": {
			opFunc:   OpFuncLE,
			n:        2,
			sets:     [][]int{{1, 2, 4, 5, 6}, {-1, 0, 3, 5}, {1, 6, 7, 8, 9, 10}},
			expected: []int{-1, 0, 2, 3, 4, 7, 8, 9, 10},
		},
	} {
		t.Run(name, func(t *testing.T) {
			actual := PerformOperation(tc.opFunc, tc.n, tc.sets...)

			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestPerformOperationEf(t *testing.T) {
	for name, tc := range map[string]struct {
		opFunc   func(idx, n uint) bool
		n        uint
		sets     []*iterableSlice
		expected []int
	}{
		"LE N=1 one set": {
			opFunc:   OpFuncLE,
			n:        1,
			sets:     []*iterableSlice{newSlice([]int{1, 2, 3})},
			expected: []int{},
		},
		"LE N=1 two sets": {
			opFunc: OpFuncLE,
			n:      1,
			sets: []*iterableSlice{
				newSlice([]int{1, 2, 3}),
				newSlice([]int{2, 3, 4}),
			},
			expected: []int{},
		},
		"LE N>1 one set": {
			opFunc:   OpFuncLE,
			n:        3,
			sets:     []*iterableSlice{newSlice([]int{1, 2, 3})},
			expected: []int{1, 2, 3},
		},
		"LE N=2 two sets equal size": {
			opFunc: OpFuncLE,
			n:      2,
			sets: []*iterableSlice{
				newSlice([]int{1, 2, 3}),
				newSlice([]int{2, 3, 4}),
			},
			expected: []int{1, 4},
		},
		"LE N=2 two sets small size=1": {
			opFunc: OpFuncLE,
			n:      2,
			sets: []*iterableSlice{
				newSlice([]int{1}),
				newSlice([]int{2}),
			},
			expected: []int{1, 2},
		},
		"LE N=2 two sets different size": {
			opFunc: OpFuncLE,
			n:      2,
			sets: []*iterableSlice{
				newSlice([]int{1}),
				newSlice([]int{2, 3}),
			},
			expected: []int{1, 2, 3},
		},
		"LE N=2 tree sets": {
			opFunc: OpFuncLE,
			n:      2,
			sets: []*iterableSlice{
				newSlice([]int{1, 2, 4, 5, 6}),
				newSlice([]int{-1, 0, 3, 5}),
				newSlice([]int{1, 6, 7, 8, 9, 10}),
			},
			expected: []int{-1, 0, 2, 3, 4, 7, 8, 9, 10},
		},
	} {
		t.Run(name, func(t *testing.T) {
			actual := PerformOperationEf(tc.opFunc, tc.n, tc.sets...)

			assert.Equal(t, tc.expected, actual)
		})
	}
}
