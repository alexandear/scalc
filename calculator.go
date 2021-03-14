package main

import (
	"container/heap"
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

type ValueIdx struct {
	Value int
	Idx   int
}

// KeyValueHeap is a min-heap of ValueIdx values comparing by Value.
type ValueIdxHeap []ValueIdx

func (h ValueIdxHeap) Len() int { return len(h) }

func (h ValueIdxHeap) Less(i, j int) bool { return h[i].Value < h[j].Value }

func (h ValueIdxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *ValueIdxHeap) Push(x interface{}) {
	*h = append(*h, x.(ValueIdx))
}

func (h *ValueIdxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}

func PerformOperationEf(opFunc func(cnt, n uint) bool, n uint, sets ...*iterableSlice) []int {
	minValueIdxes := make([]ValueIdx, 0, len(sets))
	for i := range sets {
		minValueIdxes = append(minValueIdxes, ValueIdx{Idx: i})
	}

	result := make([]int, 0)
	var valueIdxHeap ValueIdxHeap

	heap.Init(&valueIdxHeap)

	for len(minValueIdxes) != 0 {
		for _, p := range minValueIdxes {
			v, ok := sets[p.Idx].Next()
			if !ok {
				continue
			}

			heap.Push(&valueIdxHeap, ValueIdx{
				Value: v,
				Idx:   p.Idx,
			})
		}

		if valueIdxHeap.Len() == 0 {
			return result
		}

		min := valueIdxHeap[0]

		var cnt uint

		for ; valueIdxHeap.Len() != 0 && valueIdxHeap[0].Value == min.Value; cnt++ {
			minValueIdxes = append(minValueIdxes, heap.Pop(&valueIdxHeap).(ValueIdx))
		}

		if opFunc(cnt, n) {
			result = append(result, min.Value)
		}
	}

	return result
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
