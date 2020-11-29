package btree

import "log"

type Node struct {
	key                 interface{}
	value               interface{}
	left, right, parent IBinaryTreeNode
}

func NewNodeInterface(key interface{}, value interface{}) *Node {
	return &Node{key: key, value: value}
}

func (n *Node) Key() interface{} {
	return n.key
}

func (n *Node) Value() interface{} {
	return n.value
}

func (n *Node) Left() IBinaryTreeNode {
	return n.left
}

func (n *Node) Right() IBinaryTreeNode {
	return n.right
}

func (n *Node) Parent() IBinaryTreeNode {
	return n.parent
}

func (n *Node) SetLeft(node IBinaryTreeNode) {
	n.left = node
}

func (n *Node) SetRight(node IBinaryTreeNode) {
	n.right = node
}

func (n *Node) SetParent(node IBinaryTreeNode) {
	n.parent = node
}

func (n *Node) Visit() {
	log.Print(n.value)
}
