package scalc

type IterableSlice struct {
	idx int
	s   []int
}

func NewIterableSlice(s []int) *IterableSlice {
	if s == nil {
		s = []int{}
	}

	return &IterableSlice{-1, s}
}

func (s *IterableSlice) Next() (value int, ok bool) {
	s.idx++

	if s.idx >= len(s.s) {
		return 0, false
	}

	return s.s[s.idx], true
}
