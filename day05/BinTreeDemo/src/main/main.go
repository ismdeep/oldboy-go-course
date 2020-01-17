package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	var tree = NewBinTree()
	for i := 0; i < 10; i++ {
		tree.InsertByRand()
	}
	tree.DFSTraverse(func(node *Node) {
		fmt.Printf("node: %p    node.left: %p    node.right: %p => %d\n", node, node.Left, node.Right, node.Val)
	})
}
