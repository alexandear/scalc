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

	valueIdxMin := &ValueIdxHeap{}
	heap.Init(valueIdxMin)

	for i, set := range sets {
		v, ok := set.Next()
		if !ok {
			continue
		}

		heap.Push(valueIdxMin, ValueIdx{
			Value: v,
			Idx:   i,
		})
	}

	smallest := (*valueIdxMin)[0]

	popped := []ValueIdx{heap.Pop(valueIdxMin).(ValueIdx)}
	cnt := uint(1)

	for valueIdxMin.Len() != 0 && (*valueIdxMin)[0].Value == smallest.Value {
		popped = append(popped, heap.Pop(valueIdxMin).(ValueIdx))

		cnt++
	}

	res := make([]int, 0)

	if opFunc(cnt, n) {
		res = append(res, smallest.Value)
	}

	if len(popped) == 0 {
		return res
	}

	for {
		for _, p := range popped {
			v, ok := sets[p.Idx].Next()
			if !ok {
				continue
			}

			heap.Push(valueIdxMin, ValueIdx{
				Value: v,
				Idx:   p.Idx,
			})
		}

		if valueIdxMin.Len() == 0 {
			return res
		}

		smallest = (*valueIdxMin)[0]

		popped = []ValueIdx{heap.Pop(valueIdxMin).(ValueIdx)}
		cnt = uint(1)

		for valueIdxMin.Len() != 0 && (*valueIdxMin)[0].Value == smallest.Value {
			popped = append(popped, heap.Pop(valueIdxMin).(ValueIdx))

			cnt++
		}

		if opFunc(cnt, n) {
			res = append(res, smallest.Value)
		}

		if len(popped) == 0 {
			return res
		}
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
