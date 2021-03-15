package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/alexandear/scalc/pkg/scalc"
)

func Test_evaluate(t *testing.T) {
	t.Run("predefined 1 expression", func(t *testing.T) {
		calculator := &Calculator{}
		defer calculator.Close()

		actual, err := calculator.evaluate(&scalc.Expression{
			Operator: scalc.OpGR,
			N:        1,
			Sets: []*scalc.Set{
				{File: newFile("c.txt")},
				{SubExpression: &scalc.Expression{
					Operator: scalc.OpEQ,
					N:        3,
					Sets: []*scalc.Set{
						{File: newFile("a.txt")},
						{File: newFile("a.txt")},
						{File: newFile("b.txt")},
					},
				}},
			},
		})

		require.NoError(t, err)
		assert.Equal(t, scalc.NewIterableSlice([]int{2, 3}), actual)
	})

	t.Run("predefined 2 expression", func(t *testing.T) {
		calculator := &Calculator{}
		defer calculator.Close()

		actual, err := calculator.evaluate(&scalc.Expression{
			Operator: scalc.OpLE,
			N:        2,
			Sets: []*scalc.Set{
				{File: newFile("a.txt")},
				{SubExpression: &scalc.Expression{
					Operator: scalc.OpGR,
					N:        1,
					Sets: []*scalc.Set{
						{File: newFile("b.txt")},
						{File: newFile("c.txt")},
					},
				}},
			},
		})

		require.NoError(t, err)
		assert.Equal(t, scalc.NewIterableSlice([]int{1, 4}), actual)
	})
}

func newFile(name string) *string {
	return &name
}
