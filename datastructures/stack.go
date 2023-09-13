package datastructures

type Stack[T any] struct {
	head *Node[T]
	Len  int
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		head: nil,
	}
}

func (s *Stack[T]) Push(v T) {
	newNode := &Node[T]{
		Val:  v,
		Next: s.head,
	}
	s.Len++
	s.head = newNode
}

func (s *Stack[T]) Pop() T {
	if s.head == nil {
		var res T
		return res
	}
	val := s.head.Val

	s.head = s.head.Next
	s.Len--
	return val
}
