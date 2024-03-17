package stack

type node[V comparable] struct {
	value V
	prev  *node[V]
}

type Stack[V comparable] struct {
	top    *node[V]
	length int
}

func New[V comparable]() *Stack[V] {
	return &Stack[V]{nil, 0}
}

func (s *Stack[V]) Push(v V) {
	n := &node[V]{v, s.top}
	s.top = n
	s.length++
}

func (s *Stack[V]) Pop() V {
	n := s.top
	s.top = n.prev
	s.length--
	return n.value

}

func (s *Stack[V]) Peek() V {
	return s.top.value
}

func (s *Stack[V]) IsEmpty() bool {
	return s.length == 0
}

func (s *Stack[V]) Len() int {
	return s.length
}
