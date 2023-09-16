package main

import (
	"github.com/Arch-4ng3l/DSA/datastructures"
)

func main() {
	linkedList := datastructures.NewLinkedList[int]()
	linkedList.Prepand(10)
	linkedList.Append(20)
	linkedList.Print()
	linkedList.Prepand(30)
	linkedList.Print()
	linkedList.Append(40)
	linkedList.Print()
	linkedList.InsertAt(5, 3)
	linkedList.InsertAt(6, 3)
	linkedList.InsertAt(7, 3)
	linkedList.InsertAt(8, 3)
	linkedList.Print()
	linkedList.RemoveAt(3)
	linkedList.Print()
	linkedList.Remove(20)
	linkedList.Print()

}
