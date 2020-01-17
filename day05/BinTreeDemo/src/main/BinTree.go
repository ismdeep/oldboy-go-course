package main

import (
	"math/rand"
)

type BinTree struct {
	Root *Node
}

/* 二叉树构造函数 */
func NewBinTree() *BinTree {
	var ans *BinTree = &BinTree{Root: nil}
	return ans
}

/* 随机在树中插入一个节点 */
func (tree *BinTree) InsertByRand() {
	if tree.Root == nil {
		tree.Root = &Node{
			Val:   rand.Intn(100),
			Left:  nil,
			Right: nil,
		}
		return
	}
	tree.Root.InsertByRand()
}

/* BinTree DFS 遍历 */
func (tree *BinTree) DFSTraverse(Op func(node *Node)) {
	tree.Root.DFSTraverse(Op)
}
