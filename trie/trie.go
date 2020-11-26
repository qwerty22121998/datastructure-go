package trie

type Trie struct {
	root *Node
}

func NewTrie() *Trie {
	return &Trie{root: &Node{
		child: make(map[rune]*Node),
		value: nil,
		end:   0,
		cnt:   0,
	}}
}

func (t *Trie) Insert(s string, value ...interface{}) {
	root := t.root
	for _, r := range s {
		root.cnt++
		if root.child[r] == nil {
			root.newNode(r)
		}
		root = root.child[r]
	}
	root.cnt++
	root.end++
	if len(value) > 0 {
		root.value = append(root.value, value...)
	}
}

func (t *Trie) Delete(s string) bool {
	runes := []rune(s)
	if len(runes) == 0 {
		return false
	}
	prefix := runes[:len(runes)-1]
	suffix := runes[len(runes)-1]
	node := t.root.findString(prefix)
	if node == nil {
		return false
	}
	child := node.child[suffix]
	if child == nil || child.end == 0 {
		return false
	}
	child.end--
	child.cnt--
	child.value = nil
	if child.cnt == 0 && len(child.child) == 0 {
		delete(node.child, suffix)
	}
	return true
}

func (t *Trie) Find(s string) (int, []interface{}) {
	node := t.root.findString([]rune(s))
	if node == nil {
		return 0, nil
	}
	return node.end, node.value
}

func (t *Trie) BeginWith(prefix string) int {
	node := t.root.findString([]rune(prefix))
	if node == nil {
		return 0
	}
	return node.cnt
}
