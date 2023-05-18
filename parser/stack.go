package parser

type sizedStack[T any] struct {
	size     int
	elements []T
}

func newStack[T any](size int) *sizedStack[T] {
	return &sizedStack[T]{
		size: 0, elements: make([]T, size),
	}
}

func (s *sizedStack[T]) resize(cap int) {
	elements := make([]T, cap)
	for i := 0; i < s.size; i++ {
		elements[i] = s.elements[i]
	}
	s.elements = elements
}

func (s *sizedStack[T]) Push(ele T) {
	if s.size == len(s.elements) {
		s.resize(2 * len(s.elements))
	}
	s.elements[s.size] = ele
	s.size++
}

func (s *sizedStack[T]) Pop() (ele T) {
	if s.IsEmpty() {
		return
	}
	s.size--
	ele = s.elements[s.size]
	if s.size > 0 && s.size == len(s.elements)/4 {
		s.resize(len(s.elements) / 2)
	}
	return ele
}

func (s *sizedStack[T]) Peep() (ele T) {
	if s.IsEmpty() {
		return
	}
	return s.elements[s.size-1]
}

func (s *sizedStack[T]) IsEmpty() bool {
	return s.size == 0
}

func (s *sizedStack[T]) Length() int {
	return s.size
}
