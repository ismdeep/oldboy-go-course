package main

import "fmt"

func main() {
	var list *LinkedList = NewLinkedList()
	list.InsertFront(5)
	list.InsertTail(1)
	list.InsertTail(2)
	list.InsertTail(4)
	list.InsertFront(3)
	list.Traverse(func(node *Node) {
		fmt.Printf("%d ", node.Val)
	})
	fmt.Printf("\n")
}
