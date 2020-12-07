package array_queue

import "github.com/qwerty22121998/datastructure-go/queue"

type ArrayQueue []interface{}

func (a *ArrayQueue) Enqueue(i interface{}) error {
	*a = append(*a, i)
	return nil
}

func (a *ArrayQueue) Dequeue() (interface{}, error) {
	if a.IsEmpty() {
		return nil, queue.EmptyErr
	}
	defer func() { *a = (*a)[1:] }()
	return (*a)[0], nil
}

func (a *ArrayQueue) IsEmpty() bool {
	return len(*a) == 0
}

func (a *ArrayQueue) Front() (interface{}, error) {
	if a.IsEmpty() {
		return nil, queue.EmptyErr
	}
	return (*a)[0], nil
}
