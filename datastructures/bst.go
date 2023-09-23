package datastructures

import "fmt"

type TreeNode[T int] struct {
	Val   T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

type Bst[T int] struct {
	Head *TreeNode[T]
}

func NewBst[T int]() *Bst[T] {
	return &Bst[T]{
		Head: nil,
	}
}

func (t *TreeNode[T]) PrintInOrder() {
	if t.Left != nil {
		t.Left.PrintInOrder()
	}
	fmt.Print(t.Val)
	fmt.Print(" ")
	if t.Right != nil {
		t.Right.PrintInOrder()
	}
}

func (t *TreeNode[T]) PrintPostOrder() {

	if t.Left != nil {
		t.Left.PrintPostOrder()
	}
	if t.Right != nil {
		t.Right.PrintPostOrder()
	}

	fmt.Print(t.Val)
	fmt.Print(" ")
}

func (t *TreeNode[T]) PrintPreOrder() {
	fmt.Print(t.Val)
	fmt.Print(" ")
	if t.Left != nil {
		t.Left.PrintPreOrder()
	}
	if t.Right != nil {
		t.Right.PrintPreOrder()
	}
}

func (t *TreeNode[T]) Push(v T) {
	if t == nil {
		t = &TreeNode[T]{Val: v}
		return
	}

	if v < t.Val {
		if t.Left == nil {
			t.Left = &TreeNode[T]{Val: v}
		} else {
			t.Left.Push(v)
		}
	} else {
		if t.Right == nil {
			t.Right = &TreeNode[T]{Val: v}
		} else {
			t.Right.Push(v)
		}
	}

}

func (t *TreeNode[T]) Contains(v T) bool {
	if t == nil {
		return false
	}

	if v < t.Val {

		return t.Left.Contains(v)
	} else if v > t.Val {
		return t.Right.Contains(v)
	}

	return true
}
func (t *TreeNode[T]) min() *TreeNode[T] {
	temp := t
	for nil != temp && temp.Left != nil {
		temp = temp.Left
	}
	return temp
}

func (t *TreeNode[T]) Pop(v T) *TreeNode[T] {
	if t.Val == v {
		if t.Left == nil && t.Right == nil {
			t = nil
			return t
		} else if t.Right == nil {
			temp := t.Left
			t = temp
			return t

		} else if t.Left == nil {
			temp := t.Right
			t = temp
			return t
		}
		temp := t.Right.min()
		t.Val = temp.Val
		t.Right = t.Right.Pop(temp.Val)

	} else if v < t.Val {
		t.Left = t.Left.Pop(v)
	} else {
		t.Right = t.Right.Pop(v)
	}

	return t

}

func (t *TreeNode[T]) IsEqual(node *TreeNode[T]) bool {
	if t == nil && node == nil {
		return true
	}

	if t == nil || node == nil {
		return false
	}

	if t.Val != node.Val {
		return false
	}

	return t.Left.IsEqual(node.Left) && t.Right.IsEqual(node.Right)

}

func (b *Bst[T]) Pop(v T) {
	if b.Head == nil {
		return
	}
	b.Head.Pop(v)
}

func (b *Bst[T]) Contains(v T) bool {
	if b.Head == nil {
		return false
	}
	return b.Head.Contains(v)

}

func (b *Bst[T]) BreadthFirstSearch(v T) bool {
	q := NewQueue[*TreeNode[T]]()
	if b.Head == nil {
		return false
	}
	q.Enque(b.Head)
	for q.Len != 0 {
		cur := q.Deque()

		if cur.Val == v {
			return true
		}

		if cur.Left != nil {
			q.Enque(cur.Left)
		}
		if cur.Right != nil {
			q.Enque(cur.Right)
		}
	}
	return false

}

func (b *Bst[T]) IsEqual(b2 *Bst[T]) bool {
	return b.Head.IsEqual(b2.Head)
}

func (b *Bst[T]) PrintInOrder() {
	if b.Head == nil {
		return
	}
	b.Head.PrintInOrder()
	fmt.Println("")
}

func (b *Bst[T]) PrintPostOrder() {
	if b.Head == nil {
		return
	}
	b.Head.PrintPostOrder()
	fmt.Println("")
}

func (b *Bst[T]) PrintPreOrder() {
	if b.Head == nil {
		return
	}
	b.Head.PrintPreOrder()
	fmt.Println("")
}

func (b *Bst[T]) Push(v T) {
	if b.Head == nil {
		b.Head = &TreeNode[T]{Val: v}
		return
	}

	b.Head.Push(v)
}
