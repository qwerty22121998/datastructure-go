package elements

type IBinaryTreeNode interface {
	Key() interface{}
	Value() interface{}
	Left() IBinaryTreeNode
	Right() IBinaryTreeNode
	Parent() IBinaryTreeNode
	Visit()
}

func InOrder(node IBinaryTreeNode) {
	if node == nil {
		return
	}
	InOrder(node.Left())
	node.Visit()
	InOrder(node.Right())
}

func PreOrder(node IBinaryTreeNode) {
	if node == nil {
		return
	}
	node.Visit()
	PreOrder(node.Left())
	PreOrder(node.Right())
}

func PosOrder(node IBinaryTreeNode) {
	if node == nil {
		return
	}
	PosOrder(node.Left())
	PosOrder(node.Right())
	node.Visit()
}
