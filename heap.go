package main

// PairHeap is a min-heap of Pair values comparing by Pair.Value.
type PairHeap []Pair

func (h PairHeap) Len() int { return len(h) }

func (h PairHeap) Less(i, j int) bool { return h[i].Value < h[j].Value }

func (h PairHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *PairHeap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}

func (h *PairHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}
