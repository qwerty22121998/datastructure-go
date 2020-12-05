package disjoint_set

import "github.com/qwerty22121998/datastructure-go/tree"

type Node tree.BinaryNode

func (n *Node) Root() *Node {
	if n.Parent != nil {
		n.Parent = (*tree.BinaryNode)((*Node)(n.Parent).Root()) //Node(n.Parent).Root()
		return (*Node)(n.Parent)
	}
	return n
}

func (n *Node) Join(b *Node) {
	n.Parent = (*tree.BinaryNode)(b.Root())
}
