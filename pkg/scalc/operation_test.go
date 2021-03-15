package scalc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculate(t *testing.T) {
	testCalculate(t, calculate)
}

func TestCalculateInefficient(t *testing.T) {
	testCalculate(t, calculateInefficient)
}

func testCalculate(t *testing.T, performFunc func(opFunc opFunc, n uint, sets []Iterator) Iterator) {
	for name, tc := range map[string]struct {
		opFunc   opFunc
		n        uint
		sets     []Iterator
		expected Iterator
	}{
		"LE N=1 one set": {
			opFunc:   opFuncLE,
			n:        1,
			sets:     []Iterator{NewIterableSlice([]int{1, 2, 3})},
			expected: NewIterableSlice(nil),
		},
		"LE N=1 two sets": {
			opFunc: opFuncLE,
			n:      1,
			sets: []Iterator{
				NewIterableSlice([]int{1, 2, 3}),
				NewIterableSlice([]int{2, 3, 4}),
			},
			expected: NewIterableSlice(nil),
		},
		"LE N>1 one set": {
			opFunc:   opFuncLE,
			n:        3,
			sets:     []Iterator{NewIterableSlice([]int{1, 2, 3})},
			expected: NewIterableSlice([]int{1, 2, 3}),
		},
		"LE N=2 two sets equal size": {
			opFunc: opFuncLE,
			n:      2,
			sets: []Iterator{
				NewIterableSlice([]int{1, 2, 3}),
				NewIterableSlice([]int{2, 3, 4}),
			},
			expected: NewIterableSlice([]int{1, 4}),
		},
		"LE N=2 two sets small size=1": {
			opFunc: opFuncLE,
			n:      2,
			sets: []Iterator{
				NewIterableSlice([]int{1}),
				NewIterableSlice([]int{2}),
			},
			expected: NewIterableSlice([]int{1, 2}),
		},
		"LE N=2 two sets different size": {
			opFunc: opFuncLE,
			n:      2,
			sets: []Iterator{
				NewIterableSlice([]int{1}),
				NewIterableSlice([]int{2, 3}),
			},
			expected: NewIterableSlice([]int{1, 2, 3}),
		},
		"LE N=2 three sets": {
			opFunc: opFuncLE,
			n:      2,
			sets: []Iterator{
				NewIterableSlice([]int{1, 2, 4, 5, 6}),
				NewIterableSlice([]int{-1, 0, 3, 5}),
				NewIterableSlice([]int{1, 6, 7, 8, 9, 10}),
			},
			expected: NewIterableSlice([]int{-1, 0, 2, 3, 4, 7, 8, 9, 10}),
		},
		"GR N=1 two sets first has less size than a second": {
			opFunc: opFuncGR,
			n:      1,
			sets: []Iterator{
				NewIterableSlice([]int{2, 3, 4}),
				NewIterableSlice([]int{1, 2, 3, 4, 5}),
			},
			expected: NewIterableSlice([]int{2, 3, 4}),
		},
		"GR N=1 two sets second has less size than a first": {
			opFunc: opFuncGR,
			n:      1,
			sets: []Iterator{
				NewIterableSlice([]int{1, 2, 3, 4, 5}),
				NewIterableSlice([]int{2, 3}),
			},
			expected: NewIterableSlice([]int{2, 3}),
		},
		"GR N=1 three sets": {
			opFunc: opFuncGR,
			n:      1,
			sets: []Iterator{
				NewIterableSlice([]int{1, 2, 4, 5, 6}),
				NewIterableSlice([]int{-1, 0, 3, 5}),
				NewIterableSlice([]int{1, 6, 7, 8, 9, 10}),
			},
			expected: NewIterableSlice([]int{1, 5, 6}),
		},
		"EQ N=1 one set": {
			opFunc:   opFuncEQ,
			n:        1,
			sets:     []Iterator{NewIterableSlice([]int{1, 2, 3})},
			expected: NewIterableSlice([]int{1, 2, 3}),
		},
		"EQ N=3 three sets": {
			opFunc: opFuncEQ,
			n:      3,
			sets: []Iterator{
				NewIterableSlice([]int{1, 2, 3}),
				NewIterableSlice([]int{1, 2, 3}),
				NewIterableSlice([]int{2, 3}),
			},
			expected: NewIterableSlice([]int{2, 3}),
		},
		"EQ N=1 three sets": {
			opFunc: opFuncEQ,
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
