package main

import "math/rand"

/* Node 构造函数 */
type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

/* Node 构造函数 */
func NewNode(val int) *Node {
	return &Node{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}

/* 节点随机插入 */
func (node *Node) InsertByRand() {
	if node == nil {
		return
	}
	flag := rand.Intn(2)
	if flag == 0 {
		if node.Left == nil {
			node.Left = NewNode(rand.Intn(100))
		} else {
			node.Left.InsertByRand()
		}
	} else {
		if node.Right == nil {
			node.Right = NewNode(rand.Intn(100))
		} else {
			node.Right.InsertByRand()
		}
	}
}

/* Node DFS 遍历 */
func (node *Node) DFSTraverse(Op func(node *Node)) {
	if node == nil {
		return
	}
	Op(node)
	node.Left.DFSTraverse(Op)
	node.Right.DFSTraverse(Op)
}

