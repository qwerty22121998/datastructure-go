package arr_stack

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/qwerty22121998/datastructure-go/queue"
)

func TestArrStack_Push(t *testing.T) {
	var s ArrStack
	s.Push(1)
	s.Push(2)
}

func TestArrStack_Pop(t *testing.T) {
	var s ArrStack
	_, err := s.Pop()
	assert.Equal(t, queue.EmptyErr, err)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	p, err := s.Pop()
	assert.Nil(t, err)
	assert.Equal(t, 3, p)
}

func TestArrStack_Top(t *testing.T) {
	var s ArrStack
	s.Push(1)
	s.Push(2)
	v, err := s.Top()
	assert.Nil(t, err)
	assert.Equal(t, 2, v)
}

func TestArrStack_IsEmpty(t *testing.T) {
	var s ArrStack
	assert.True(t, s.IsEmpty())
	s.Push(10)
	assert.False(t, s.IsEmpty())
	s.Pop()
	assert.True(t, s.IsEmpty())

}
