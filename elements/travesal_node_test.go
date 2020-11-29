package elements

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

type TmpNode struct {
	key   int
	value int
	left  INode
	right INode
}

func (t *TmpNode) Key() interface{} {
	return t.key
}

func (t *TmpNode) Value() interface{} {
	return t.value
}

func (t *TmpNode) Left() INode {
	return t.left
}

func (t *TmpNode) Right() INode {
	return t.right
}

func (t *TmpNode) Print() {
	log.Print(t.value)
}

var root *TmpNode

func init() {
	log.SetFlags(0)
	root = NewNode(1)
	n2 := NewNode(2)
	n3 := NewNode(3)
	n4 := NewNode(4)
	n5 := NewNode(5)
	n2.left = n4
	n2.right = n5
	root.left = n2
	root.right = n3
}

func captureOutput(f func()) string {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	f()
	log.SetOutput(os.Stderr)
	return buf.String()
}

func NewNode(k int) *TmpNode {
	return &TmpNode{
		key:   k,
		value: k,
	}
}

func TestInOrder(t *testing.T) {
	expected := "4\n2\n5\n1\n3\n"
	assert.Equal(t, captureOutput(func() {
		InOrder(root)
	}),expected)
}

func TestPreOrder(t *testing.T) {
	expected := "1\n2\n4\n5\n3\n"
	assert.Equal(t, captureOutput(func() {
		PreOrder(root)
	}),expected)
}

func TestPosOrder(t *testing.T) {
	expected := "4\n5\n2\n3\n1\n"
	assert.Equal(t, captureOutput(func() {
		PosOrder(root)
	}),expected)
}
