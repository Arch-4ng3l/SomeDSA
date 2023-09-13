package datastructures

type Node[T any] struct {
	Val  T
	Next *Node[T]
}

type Queue[T any] struct {
	head *Node[T]
	tail *Node[T]
	Len  int
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		head: nil,
		tail: nil,
	}
}

func (q *Queue[T]) Enque(v T) {
	q.Len++
	if q.head == nil {
		q.head = &Node[T]{
			Val:  v,
			Next: nil,
		}
		return
	}
	if q.tail == nil {
		q.tail = &Node[T]{
			Val:  v,
			Next: nil,
		}
		q.head.Next = q.tail
		return

	}
	q.tail.Next = &Node[T]{
		Val:  v,
		Next: nil,
	}
	q.tail = q.tail.Next
}

func (q *Queue[T]) Deque() T {
	if q.head == nil {
		var res T
		return res
	}
	val := q.head.Val
	q.head = q.head.Next
	q.Len--

	return val
}
