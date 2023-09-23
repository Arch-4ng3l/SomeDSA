package main

import (
	"fmt"

	"github.com/Arch-4ng3l/DSA/datastructures"
)

func main() {
	bst := datastructures.NewBst[int]()
	bst.Push(10)
	bst.Push(12)
	bst.Push(11)
	bst.Push(9)
	bst.Push(8)

	bst2 := datastructures.NewBst[int]()
	bst2.Push(10)
	bst2.Push(11)
	bst2.Push(8)
	bst2.Push(12)
	bst2.Push(9)

	fmt.Println(bst.IsEqual(bst2))

}
