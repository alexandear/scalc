package main

import (
	"sort"
)

func PerformOperation(opFunc func(idx, n uint) bool, n uint, sets ...[]int) []int {
	counts := make(map[int]uint)

	for _, set := range sets {
		for _, number := range set {
			counts[number]++
		}
	}

	var res []int
	for number, c := range counts {
		if opFunc(c, n) {
			res = append(res, number)
		}
	}

	sort.Ints(res)

	return res
}

func OpFuncEQ(idx, n uint) bool {
	return idx == n
}

func OpFuncLE(idx, n uint) bool {
	return idx < n
}

func OpFuncGR(idx, n uint) bool {
	return idx > n
}
