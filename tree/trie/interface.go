package trie

// ITrie : Trie interface
type ITrie interface {
	Insert(s string, value ...interface{})
	Delete(s string) bool
	Find(s string) (int, []interface{})
	BeginWith(prefix string) int
}
