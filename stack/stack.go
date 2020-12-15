package stack

type IStack interface {
	Push(i interface{})
	Top() (interface{}, error)
	Pop() (interface{}, error)
	IsEmpty() bool
}
