package elements

type INode interface {
	Key() interface{}
	Value() interface{}
	Left() INode
	Right() INode
	Visit()
}

func InOrder(node INode) {
	if node == nil {
		return
	}
	InOrder(node.Left())
	node.Visit()
	InOrder(node.Right())
}

func PreOrder(node INode) {
	if node == nil {
		return
	}
	node.Visit()
	PreOrder(node.Left())
	PreOrder(node.Right())
}

func PosOrder(node INode) {
	if node == nil {
		return
	}
	PosOrder(node.Left())
	PosOrder(node.Right())
	node.Visit()
}
