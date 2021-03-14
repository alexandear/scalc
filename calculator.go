package main

import (
	"math"
	"sort"
)

func PerformOperation(opFunc func(cnt, n uint) bool, n uint, sets ...[]int) []int {
	counts := make(map[int]uint)

	for _, set := range sets {
		for _, number := range set {
			counts[number]++
		}
	}

	res := make([]int, 0)

	for number, cnt := range counts {
		if opFunc(cnt, n) {
			res = append(res, number)
		}
	}

	sort.Ints(res)

	return res
}

type intIterator interface {
	Next() (value int, ok bool)
}

type iterableSlice struct {
	idx int
	s   []int
}

func (s *iterableSlice) Next() (value int, ok bool) {
	s.idx++

	if s.idx >= len(s.s) {
		return 0, false
	}

	return s.s[s.idx], true
}

func newSlice(s []int) *iterableSlice {
	return &iterableSlice{-1, s}
}

func PerformOperationEf(opFunc func(cnt, n uint) bool, n uint, sets ...*iterableSlice) []int {
	if n == 1 {
		return []int{}
	}

	if len(sets) == 1 {
		var res []int

		for {
			v, ok := sets[0].Next()
			if !ok {
				break
			}

			res = append(res, v)
		}

		return res
	}

	res := make([]int, 0)

	line := make([]int, len(sets))
	exist := make([]bool, len(sets))

	for i, set := range sets {
		v, ok := set.Next()
		if !ok {
			continue
		}

		line[i] = v
		exist[i] = true
	}

	smallest := math.MaxInt32

	for _, l := range line {
		if l < smallest {
			smallest = l
		}
	}

	var cnt uint

	for _, l := range line {
		if l == smallest {
			cnt++
		}
	}

	if cnt < n {
		res = append(res, smallest)
	}

	idx := make([]bool, len(sets))

	for i, l := range line {
		if l == smallest {
			idx[i] = true
		}
	}

	atLeastOnceExist := false

	for _, e := range exist {
		if e {
			atLeastOnceExist = true

			break
		}
	}

	if !atLeastOnceExist {
		return res
	}

	for {
		for i, id := range idx {
			if !id {
				continue
			}

			v, ok := sets[i].Next()
			if !ok {
				exist[i] = false

				continue
			}

			line[i] = v
		}

		smallest = math.MaxInt32
		for i, l := range line {
			if l < smallest {
				if exist[i] {
					smallest = l
				}
			}
		}

		if smallest == math.MaxInt32 {
			return res
		}

		cnt = 0

		for _, l := range line {
			if l == smallest {
				cnt++
			}
		}

		if opFunc(cnt, n) {
			res = append(res, smallest)
		}

		idx = make([]bool, len(sets))

		for i, l := range line {
			if l == smallest {
				idx[i] = true
			}
		}

		atLeastOnceExist = false

		for _, e := range exist {
			if e {
				atLeastOnceExist = true

				break
			}
		}

		if atLeastOnceExist {
			continue
		}

		return res
	}
}

func OpFuncEQ(cnt, n uint) bool {
	return cnt == n
}

func OpFuncLE(cnt, n uint) bool {
	return cnt < n
}

func OpFuncGR(cnt, n uint) bool {
	return cnt > n
}
