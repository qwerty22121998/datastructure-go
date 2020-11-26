package trie

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestTrie_Insert(t *testing.T) {
	trie := NewTrie()
	trie.Insert("abc")
}

func TestTrie_Find(t *testing.T) {
	trie := NewTrie()
	trie.Insert("abc", 1, 2, 3)
	expectedCnt, expectedValue := 1, []interface{}{1, 2, 3}
	cnt, value := trie.Find("abc")
	assert.Equal(t, expectedCnt, cnt)
	assert.Equal(t, expectedValue, value)
}

func TestTrie_Delete(t *testing.T) {
	trie := NewTrie()
	trie.Insert("abcd")
	trie.Insert("abc")
	log.Print(trie.root.findString([]rune("abc")))
	assert.True(t, trie.Delete("abc"))
	log.Print(trie.root.findString([]rune("abc")))
	assert.False(t, trie.Delete("abc"))
	cnt, value := trie.Find("abcd")
	assert.Equal(t, 1, cnt)
	assert.Equal(t, 0, len(value))
}

func TestTrie_BeginWith(t *testing.T) {
	trie := NewTrie()
	trie.Insert("abcd")
	assert.Equal(t, trie.BeginWith("abc"), 1)
	assert.Equal(t, trie.BeginWith("abd"), 0)
	trie.Insert("abce")
	assert.Equal(t, trie.BeginWith("abc"), 2)
	assert.Equal(t, trie.BeginWith("abcd"), 1)
}