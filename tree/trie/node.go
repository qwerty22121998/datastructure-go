package trie

type Node struct {
	parent *Node
	child  map[rune]*Node
	value  []interface{}
	end    int
	cnt    int
}

func (n *Node) newNode(r rune) *Node {
	n.child[r] = &Node{
		parent: n,
		child:  make(map[rune]*Node),
		value: nil,
		end:    0,
		cnt:    0,
	}
	return n.child[r]
}

func (n *Node) findString(s []rune) *Node {
	root := n
	for _, r := range s {
		if root.child[r] == nil {
			return nil
		}
		root = root.child[r]
	}
	return root
}