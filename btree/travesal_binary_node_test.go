package btree

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

var root *Node

func init() {
	log.SetFlags(0)
	root = NewNodeInterface(1, 1)
	n2 := NewNodeInterface(2, 2)
	n3 := NewNodeInterface(3, 3)
	n4 := NewNodeInterface(4, 4)
	n5 := NewNodeInterface(5, 5)
	n2.SetLeft(n4)
	n2.SetRight(n5)
	root.SetLeft(n2)
	root.SetRight(n3)
}

func captureOutput(f func()) string {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	f()
	log.SetOutput(os.Stderr)
	return buf.String()
}

func TestInOrder(t *testing.T) {
	expected := "4\n2\n5\n1\n3\n"
	assert.Equal(t, captureOutput(func() {
		InOrder(root)
	}), expected)
}

func TestPreOrder(t *testing.T) {
	expected := "1\n2\n4\n5\n3\n"
	assert.Equal(t, captureOutput(func() {
		PreOrder(root)
	}), expected)
}

func TestPosOrder(t *testing.T) {
	expected := "4\n5\n2\n3\n1\n"
	assert.Equal(t, captureOutput(func() {
		PosOrder(root)
	}), expected)
}
