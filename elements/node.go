package elements

type INode interface {
	Key() interface{}
	Value() interface{}
	Left() INode
	Right() INode
	Print()
}

func InOrder(node INode) {
	if node == nil {
		return
	}
	InOrder(node.Left())
	node.Print()
	InOrder(node.Right())
}

func PreOrder(node INode) {
	if node == nil {
		return
	}
	node.Print()
	PreOrder(node.Left())
	PreOrder(node.Right())
}

func PosOrder(node INode) {
	if node == nil {
		return
	}
	PosOrder(node.Left())
	PosOrder(node.Right())
	node.Print()
}
