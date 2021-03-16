// +build integration

package test

import (
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/alexandear/scalc/internal/calc"
	"github.com/alexandear/scalc/internal/parser"
)

func TestHuge(t *testing.T) {
	createHugeFile(t, "aHuge.txt", 0, 1000005)
	createHugeFile(t, "bHuge.txt", 1000000, 2000001)

	calculator := calc.NewCalculator(parser.NewParser(), calc.NewFileIterator())
	iter, err := calculator.Calculate("[ EQ 2 aHuge.txt [ EQ 2 aHuge.txt bHuge.txt ] ]")

	require.NoError(t, err)
	var actual []int
	for v, ok := iter.Next(); ok; v, ok = iter.Next() {
		actual = append(actual, v)
	}
	assert.Equal(t, []int{1000000, 1000001, 1000002, 1000003, 1000004}, actual)
}

func createHugeFile(t *testing.T, filename string, start, finish int) {
	if err := os.Remove(filename); !os.IsNotExist(err) {
		require.NoError(t, err)
	}

	file, err := os.Create(filename)
	require.NoError(t, err)
	defer file.Close()

	for i := start; i < finish; i++ {
		_, err := file.WriteString(strconv.Itoa(i) + "\n")
		require.NoError(t, err)
	}
}
