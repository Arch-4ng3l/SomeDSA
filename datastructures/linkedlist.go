package datastructures

import "fmt"

type ListNode[T comparable] struct {
	Value T
	Prev  *ListNode[T]
	Next  *ListNode[T]
}

type LinkedList[T comparable] struct {
	Length int
	head   *ListNode[T]
	tail   *ListNode[T]
}

func NewLinkedList[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{
		Length: 0,
		head:   nil,
		tail:   nil,
	}
}

func (ll *LinkedList[T]) Print() {
	cur := ll.head
	fmt.Println(ll.Length)
	for i := 0; i < ll.Length-1; i++ {
		if cur == nil {
			continue
		}
		fmt.Print(cur.Value)
		fmt.Print(" -> ")
		cur = cur.Next
	}
	fmt.Print(ll.tail.Value)
	fmt.Print("\n")
}

func (ll *LinkedList[T]) Prepand(v T) {
	newNode := &ListNode[T]{
		Value: v,
	}

	ll.Length++

	if ll.head == nil {
		ll.head = newNode
		ll.tail = newNode
		return
	}

	newNode.Next = ll.head
	ll.head.Prev = newNode
	ll.head = newNode
}

func (ll *LinkedList[T]) InsertAt(v T, idx int) {
	if idx > ll.Length {
		return
	} else if idx == ll.Length {
		ll.Append(v)
		return
	} else if idx == 0 {
		ll.Prepand(v)
		return
	}

	cur := ll.getAt(idx)

	ll.Length++
	newNode := &ListNode[T]{Value: v}

	newNode.Next = cur
	newNode.Prev = cur.Prev

	cur.Prev.Next = newNode
	cur.Prev = newNode

}

func (ll *LinkedList[T]) Append(v T) {
	ll.Length++
	newNode := &ListNode[T]{Value: v}
	if ll.tail == nil {
		ll.tail = newNode
		ll.head = newNode
		return
	}
	newNode.Prev = ll.tail
	ll.tail.Next = newNode

	ll.tail = newNode
}

func (ll *LinkedList[T]) Remove(v T) T {
	cur := ll.head
	if cur == nil {
		ll.Length--
		ll.head = nil
		val := ll.tail.Value
		ll.tail = nil
		return val
	}

	for i := 0; i < ll.Length; i++ {
		if cur == nil {
			var res T
			return res
		}
		if cur.Value == v {
			break
		}
		cur = cur.Next
	}
	if cur == nil {
		var res T
		return res
	}
	return ll.removeNode(cur)

}

func (ll *LinkedList[T]) getAt(idx int) *ListNode[T] {
	cur := ll.head
	for i := 0; i < idx; i++ {
		if cur != nil {
			cur = cur.Next
		}
	}
	return cur
}

func (ll *LinkedList[T]) RemoveAt(idx int) T {
	node := ll.getAt(idx)
	return ll.removeNode(node)

}

func (ll *LinkedList[T]) removeNode(node *ListNode[T]) T {
	ll.Length--

	if ll.Length == 0 {
		val := ll.head.Value
		ll.head = nil
		ll.tail = nil
		return val
	}

	if node.Prev != nil {
		node.Prev.Next = node.Next
	}
	if node.Next != nil {
		node.Next.Prev = node.Prev
	}

	if node == ll.head {
		ll.head = node.Next
	}

	if node == ll.tail {
		ll.tail = node.Prev
	}

	node.Next = nil
	node.Prev = nil
	return node.Value

}

func (ll *LinkedList[T]) Get(idx int) T {
	return ll.getAt(idx).Value
}
