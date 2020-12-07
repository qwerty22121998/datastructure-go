package array_queue

import (
	"github.com/qwerty22121998/datastructure-go/queue"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArrayQueue_IsEmpty(t *testing.T) {
	var q ArrayQueue
	assert.True(t, q.IsEmpty())
	q.Enqueue(1)
	assert.False(t, q.IsEmpty())
	q.Dequeue()
	assert.True(t, q.IsEmpty())
}

func TestArrayQueue_Enqueue(t *testing.T) {
	var q ArrayQueue
	if err := q.Enqueue(1); err != nil {
		t.Error(err)
	}
	if err := q.Enqueue(2); err != nil {
		t.Error(err)
	}
	if err := q.Enqueue(3); err != nil {
		t.Error(err)
	}
	assert.Equal(t, ArrayQueue{1, 2, 3}, q)
}

func TestArrayQueue_Front(t *testing.T) {
	var q ArrayQueue
	if err := q.Enqueue(1); err != nil {
		t.Error(err)
	}
	if err := q.Enqueue(2); err != nil {
		t.Error(err)
	}
	if err := q.Enqueue(3); err != nil {
		t.Error(err)
	}
	f, err := q.Front()
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 1, f)
}

func TestArrayQueue_Dequeue(t *testing.T) {
	var q ArrayQueue
	f, err := q.Dequeue()
	assert.Nil(t, f)
	assert.Equal(t, queue.EmptyErr, err)
	assert.Nil(t, q.Enqueue(1))
	assert.Nil(t, q.Enqueue(2))
	assert.Nil(t, q.Enqueue(3))
	f, err = q.Dequeue()
	assert.Nil(t, err)
	assert.Equal(t, 1, f)
	f, err = q.Dequeue()
	assert.Nil(t, err)
	assert.Equal(t, 2, f)
}
