package scalc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculate(t *testing.T) {
	for name, tc := range map[string]struct {
		operator Operator
		n        uint
		sets     []Iterator
		expected Iterator
	}{
		"LE N=1 one set": {
			operator: OpLE,
			n:        1,
			sets:     []Iterator{NewIterableSlice([]int{1, 2, 3})},
			expected: NewIterableSlice(nil),
		},
		"LE N=1 two sets": {
			operator: OpLE,
			n:        1,
			sets: []Iterator{
				NewIterableSlice([]int{1, 2, 3}),
				NewIterableSlice([]int{2, 3, 4}),
			},
			expected: NewIterableSlice(nil),
		},
		"LE N>1 one set": {
			operator: OpLE,
			n:        3,
			sets:     []Iterator{NewIterableSlice([]int{1, 2, 3})},
			expected: NewIterableSlice([]int{1, 2, 3}),
		},
		"LE N=2 two sets equal size": {
			operator: OpLE,
			n:        2,
			sets: []Iterator{
				NewIterableSlice([]int{1, 2, 3}),
				NewIterableSlice([]int{2, 3, 4}),
			},
			expected: NewIterableSlice([]int{1, 4}),
		},
		"LE N=2 two sets small size=1": {
			operator: OpLE,
			n:        2,
			sets: []Iterator{
				NewIterableSlice([]int{1}),
				NewIterableSlice([]int{2}),
			},
			expected: NewIterableSlice([]int{1, 2}),
		},
		"LE N=2 two sets different size": {
			operator: OpLE,
			n:        2,
			sets: []Iterator{
				NewIterableSlice([]int{1}),
				NewIterableSlice([]int{2, 3}),
			},
			expected: NewIterableSlice([]int{1, 2, 3}),
		},
		"LE N=2 three sets": {
			operator: OpLE,
			n:        2,
			sets: []Iterator{
				NewIterableSlice([]int{1, 2, 4, 5, 6}),
				NewIterableSlice([]int{-1, 0, 3, 5}),
				NewIterableSlice([]int{1, 6, 7, 8, 9, 10}),
			},
			expected: NewIterableSlice([]int{-1, 0, 2, 3, 4, 7, 8, 9, 10}),
		},
		"GR N=1 two sets first has less size than a second": {
			operator: OpGR,
			n:        1,
			sets: []Iterator{
				NewIterableSlice([]int{2, 3, 4}),
				NewIterableSlice([]int{1, 2, 3, 4, 5}),
			},
			expected: NewIterableSlice([]int{2, 3, 4}),
		},
		"GR N=1 two sets second has less size than a first": {
			operator: OpGR,
			n:        1,
			sets: []Iterator{
				NewIterableSlice([]int{1, 2, 3, 4, 5}),
				NewIterableSlice([]int{2, 3}),
			},
			expected: NewIterableSlice([]int{2, 3}),
		},
		"GR N=1 three sets": {
			operator: OpGR,
			n:        1,
			sets: []Iterator{
				NewIterableSlice([]int{1, 2, 4, 5, 6}),
				NewIterableSlice([]int{-1, 0, 3, 5}),
				NewIterableSlice([]int{1, 6, 7, 8, 9, 10}),
			},
			expected: NewIterableSlice([]int{1, 5, 6}),
		},
		"EQ N=1 one set": {
			operator: OpEQ,
			n:        1,
			sets:     []Iterator{NewIterableSlice([]int{1, 2, 3})},
			expected: NewIterableSlice([]int{1, 2, 3}),
		},
		"EQ N=3 three sets": {
			operator: OpEQ,
			n:        3,
			sets: []Iterator{
				NewIterableSlice([]int{1, 2, 3}),
				NewIterableSlice([]int{1, 2, 3}),
				NewIterableSlice([]int{2, 3}),
			},
			expected: NewIterableSlice([]int{2, 3}),
		},
		"EQ N=1 three sets": {
			operator: OpEQ,
			n:        1,
			sets: []Iterator{
				NewIterableSlice([]int{1, 2, 4, 5, 6}),
				NewIterableSlice([]int{-1, 0, 3, 5}),
				NewIterableSlice([]int{1, 6, 7, 8, 9, 10}),
			},
			expected: NewIterableSlice([]int{-1, 0, 2, 3, 4, 7, 8, 9, 10}),
		},
	} {
		t.Run(name, func(t *testing.T) {
			actual := Calculate(tc.operator, tc.n, tc.sets)

			assert.Equal(t, tc.expected, actual)
		})
	}
}

func BenchmarkCalculate10(b *testing.B) {
	benchmarkCalculate(10, Calculate, b)
}

func BenchmarkCalculate100(b *testing.B) {
	benchmarkCalculate(100, Calculate, b)
}

func BenchmarkCalculate500(b *testing.B) {
	benchmarkCalculate(500, Calculate, b)
}

func BenchmarkCalculate750(b *testing.B) {
	benchmarkCalculate(750, Calculate, b)
}

func BenchmarkCalculate1000(b *testing.B) {
	benchmarkCalculate(1000, Calculate, b)
}

func BenchmarkCalculate2500(b *testing.B) {
	benchmarkCalculate(2500, Calculate, b)
}

func BenchmarkCalculate5000(b *testing.B) {
	benchmarkCalculate(5000, Calculate, b)
}

func BenchmarkCalculate7500(b *testing.B) {
	benchmarkCalculate(7500, Calculate, b)
}

func BenchmarkCalculate10000(b *testing.B) {
	benchmarkCalculate(10000, Calculate, b)
}

func BenchmarkCalculate25000(b *testing.B) {
	benchmarkCalculate(25000, Calculate, b)
}

func BenchmarkCalculate50000(b *testing.B) {
	benchmarkCalculate(50000, Calculate, b)
}

func BenchmarkCalculate75000(b *testing.B) {
	benchmarkCalculate(75000, Calculate, b)
}

func BenchmarkCalculate100000(b *testing.B) {
	benchmarkCalculate(100000, Calculate, b)
}

func BenchmarkCalculate1000000(b *testing.B) {
	benchmarkCalculate(1000000, Calculate, b)
}

var res Iterator

func benchmarkCalculate(size int, calculate func(operator Operator, n uint, iters []Iterator) Iterator, b *testing.B) {
	arrA := make([]int, size)
	arrB := make([]int, size)
	for i := 0; i < size; i++ {
		arrA[i] = i
		arrB[i] = i + 10
	}

	var r Iterator

	for n := 0; n < b.N; n++ {
		r = calculate(OpEQ, 2, []Iterator{NewIterableSlice(arrA), NewIterableSlice(arrB)})
	}

	res = r
}
