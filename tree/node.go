package tree

type BinaryNode struct {
	Key                 interface{}
	Value               interface{}
	Left, Right, Parent *BinaryNode
}
