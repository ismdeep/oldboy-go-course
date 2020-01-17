package main

type LinkedList struct {
	Head *Node
	Tail *Node
}

/* 链表构造函数 */
func NewLinkedList () *LinkedList {
	var list = &LinkedList{
		Head: nil,
		Tail: nil,
	}
	list.Head = &Node{
		Val:  0,
		Next: nil,
	}
	list.Tail = list.Head
	return list
}

/* 判断链表是否为空 */
func (list *LinkedList) IsEmpty() bool {
	if list.Head.Next == nil {
		return true
	}
	return false
}

/* 在尾部插入 */
func (list *LinkedList) InsertTail(val int) {
	var node *Node = &Node{
		Val:  val,
		Next: nil,
	}
	list.Tail.Next = node
	list.Tail = list.Tail.Next
}

/* 在头部插入 */
func (list *LinkedList) InsertFront(val int) {
	var node *Node = &Node{
		Val:  val,
		Next: nil,
	}
	node.Next = list.Head.Next
	if list.Head.Next == nil {
		list.Tail = node
	}
	list.Head.Next = node
}

/* 遍历 */
func (list *LinkedList) Traverse(Op func(node *Node)) {
	var cur = list.Head
	for cur.Next != nil {
		cur = cur.Next
		Op(cur)
	}
}