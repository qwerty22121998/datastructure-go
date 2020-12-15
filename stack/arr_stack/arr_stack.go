package arr_stack

import (
	"github.com/qwerty22121998/datastructure-go/queue"
)

type ArrStack struct {
	elements []interface{}
}

func (a *ArrStack) Push(i interface{}) {
	a.elements = append(a.elements, i)
}

func (a *ArrStack) Top() (interface{}, error) {
	if a.IsEmpty() {
		return nil, queue.EmptyErr
	}
	return a.elements[len(a.elements)-1], nil
}

func (a *ArrStack) Pop() (interface{}, error) {
	result, err := a.Top()
	if err != nil {
		return nil, err
	}
	a.elements = a.elements[:len(a.elements)-1]
	return result, nil
}

func (a *ArrStack) IsEmpty() bool {
	return len(a.elements) == 0
}
