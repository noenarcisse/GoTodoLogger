package sortedmap

import (
	"cmp"
	"iter"
	"slices"
)

type SortedMap[T cmp.Ordered, U any] struct {
	inner map[T]U
	order []T
}

func NewSortedMap[T cmp.Ordered, U any]() *SortedMap[T, U] {
	return &SortedMap[T, U]{
		inner: map[T]U{},
		order: []T{},
	}
}

func (s *SortedMap[T, U]) sort() {
	slices.Sort(s.order)
}

func (s *SortedMap[T, U]) Add(key T, val U) {
	if _, ok := s.inner[key]; !ok {
		s.order = append(s.order, key)
	}
	s.inner[key] = val
}

func (s *SortedMap[T, U]) Items() iter.Seq2[T, U] {
	s.sort()
	return func(yield func(T, U) bool) {
		for _, e := range s.order {
			if !yield(e, s.inner[e]) {
				return
			}
		}
	}
}
