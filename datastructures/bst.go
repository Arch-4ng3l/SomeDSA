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

func (t *TreeNode[T]) Print() {
	if t.Left != nil {
		t.Left.Print()
	}
	fmt.Print(t.Val)
	fmt.Print(" ")
	if t.Right != nil {
		t.Right.Print()
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

func (b *Bst[T]) Print() {
	if b.Head == nil {
		return
	}
	b.Head.Print()
	fmt.Println("")
}

func (b *Bst[T]) Push(v T) {
	if b.Head == nil {
		b.Head = &TreeNode[T]{Val: v}
		return
	}

	b.Head.Push(v)
}
