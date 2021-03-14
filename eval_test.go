package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvaluate(t *testing.T) {
	t.Run("predefined 1 expression", func(t *testing.T) {
		actual := Evaluate(&Expression{
			Operator: OpGR,
			N:        1,
			Sets: []*Set{
				{File: newFile("test/c.txt")},
				{SubExpression: &Expression{
					Operator: OpEQ,
					N:        3,
					Sets: []*Set{
						{File: newFile("test/a.txt")},
						{File: newFile("test/a.txt")},
						{File: newFile("test/b.txt")},
					},
				}},
			},
		})

		assert.Equal(t, NewIterableSlice([]int{2, 3}), actual)
	})

	t.Run("predefined 2 expression", func(t *testing.T) {
		actual := Evaluate(&Expression{
			Operator: OpLE,
			N:        2,
			Sets: []*Set{
				{File: newFile("test/a.txt")},
				{SubExpression: &Expression{
					Operator: OpGR,
					N:        1,
					Sets: []*Set{
						{File: newFile("test/b.txt")},
						{File: newFile("test/c.txt")},
					},
				}},
			},
		})

		assert.Equal(t, NewIterableSlice([]int{1, 4}), actual)
	})
}
