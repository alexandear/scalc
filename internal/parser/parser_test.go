package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/alexandear/scalc/pkg/scalc"
)

func TestParse(t *testing.T) {
	t.Run("parse success", func(t *testing.T) {
		for name, tc := range map[string]struct {
			str      string
			expected *Expression
		}{
			"one set": {
				str: `[ GR 1 a.txt ]`,
				expected: &Expression{
					Operator: scalc.OpGR,
					N:        1,
					Sets: []*Set{
						{File: newFile("a.txt")},
					},
				},
			},
			"sets": {
				str: `[ EQ 2 a.txt b.txt ]`,
				expected: &Expression{
					Operator: scalc.OpEQ,
					N:        2,
					Sets: []*Set{
						{File: newFile("a.txt")},
						{File: newFile("b.txt")},
					},
				},
			},
			"sets with sub expression": {
				str: `[ EQ 2 a.txt b.txt [ LE 5 a.txt b.txt ] d.txt ]`,
				expected: &Expression{
					Operator: scalc.OpEQ,
					N:        2,
					Sets: []*Set{
						{File: newFile("a.txt")},
						{File: newFile("b.txt")},
						{SubExpression: &Expression{
							Operator: scalc.OpLE,
							N:        5,
							Sets: []*Set{
								{File: newFile("a.txt")},
								{File: newFile("b.txt")},
							},
						}},
						{File: newFile("d.txt")},
					},
				},
			},
			"long filename": {
				str: `[ EQ 2 aHuge.txt b2.txt ]`,
				expected: &Expression{
					Operator: scalc.OpEQ,
					N:        2,
					Sets: []*Set{
						{File: newFile("aHuge.txt")},
						{File: newFile("b2.txt")},
					},
				},
			},
			"spaces": {
				str: ` [GR  1 a.txt     c.txt  ]   `,
				expected: &Expression{
					Operator: scalc.OpGR,
					N:        1,
					Sets: []*Set{
						{File: newFile("a.txt")},
						{File: newFile("c.txt")},
					},
				},
			},
			"predefined 1 expression": {
				str: `[ GR 1 c.txt [ EQ 3 a.txt a.txt b.txt ] ]`,
				expected: &Expression{
					Operator: scalc.OpGR,
					N:        1,
					Sets: []*Set{
						{File: newFile("c.txt")},
						{SubExpression: &Expression{
							Operator: scalc.OpEQ,
							N:        3,
							Sets: []*Set{
								{File: newFile("a.txt")},
								{File: newFile("a.txt")},
								{File: newFile("b.txt")},
							},
						}},
					},
				},
			},
			"predefined 2 expression": {
				str: `[ LE 2 a.txt [ GR 1 b.txt c.txt ] ]`,
				expected: &Expression{
					Operator: scalc.OpLE,
					N:        2,
					Sets: []*Set{
						{File: newFile("a.txt")},
						{SubExpression: &Expression{
							Operator: scalc.OpGR,
							N:        1,
							Sets: []*Set{
								{File: newFile("b.txt")},
								{File: newFile("c.txt")},
							},
						}},
					},
				},
			},
			"complex expression": {
				str: `[ LE 3 a.txt b.txt c.txt [ GR 1 d.txt e.txt ] [ EQ 2 f.txt [ LE 1 g.txt h.txt ] ] ]`,
				expected: &Expression{
					Operator: scalc.OpLE,
					N:        3,
					Sets: []*Set{
						{File: newFile("a.txt")},
						{File: newFile("b.txt")},
						{File: newFile("c.txt")},
						{SubExpression: &Expression{
							Operator: scalc.OpGR,
							N:        1,
							Sets: []*Set{
								{File: newFile("d.txt")},
								{File: newFile("e.txt")},
							},
						}},
						{SubExpression: &Expression{
							Operator: scalc.OpEQ,
							N:        2,
							Sets: []*Set{
								{File: newFile("f.txt")},
								{SubExpression: &Expression{
									Operator: scalc.OpLE,
									N:        1,
									Sets: []*Set{
										{File: newFile("g.txt")},
										{File: newFile("h.txt")},
									},
								}},
							},
						}},
					},
				},
			},
		} {
			t.Run(name, func(t *testing.T) {
				parser := NewParser()

				actual, err := parser.Parse(tc.str)

				require.NoError(t, err)
				assert.Equal(t, tc.expected, actual)
			})
		}
	})

	t.Run("parse fails", func(t *testing.T) {
		for name, tc := range map[string]struct {
			str string
		}{
			"empty": {
				str: ` `,
			},
			"missing brackets": {
				str: `GR 2 a.txt`,
			},
			"brackets inconsistent": {
				str: `[ GR 2 a.txt`,
			},
			"missing operator": {
				str: `[ 2 a.txt ]`,
			},
			"missing N": {
				str: `[ GR a.txt ]`,
			},
			"negative N": {
				str: `[ GR -2 a.txt ]`,
			},
			"zero N": {
				str: `[ GR 0 a.txt ]`,
			},
			"wrong operator": {
				str: `[ IN 2 a.txt ]`,
			},
			"wrong filename": {
				str: `[ GR 2 atxt ]`,
			},
			"wrong sub expression": {
				str: `[ GR 2 a.txt [ LE 2 ] ]`,
			},
		} {
			t.Run(name, func(t *testing.T) {
				parser := NewParser()

				actual, err := parser.Parse(tc.str)

				assert.Error(t, err)
				assert.Nil(t, actual)
			})
		}
	})
}

func newFile(file string) *string {
	return &file
}
