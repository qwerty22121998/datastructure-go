package queue

import "errors"

var EmptyErr = errors.New("queue is empty")

type IQueue interface {
	Enqueue(i interface{}) error
	Dequeue() (interface{}, error)
	IsEmpty() bool
	Front() (interface{}, error)
}
