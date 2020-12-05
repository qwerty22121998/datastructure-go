package graph

// IGraphInt: Interface for storing graph data structure
type IGraphInt interface {
	Add(i, j int, l int64) error
	Length(i, j int) int64
	Adj(i int) []int
}
