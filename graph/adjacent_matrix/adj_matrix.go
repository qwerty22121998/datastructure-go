package adjacent_matrix

import (
	"errors"
)

type AdjMatrix struct {
	size   int
	matrix [][]int64
}

var ErrSizeLimit = errors.New("out of size")

func (a *AdjMatrix) isOutOfSize(i, j int) bool {
	return i < 0 || j < 0 || i >= a.size || j >= a.size
}

//NewAdjMatrix: Create new adjacent matrix with defined size
func NewAdjMatrix(size int) *AdjMatrix {
	m := make([][]int64, size)
	for i := range m {
		m[i] = make([]int64, size)
	}
	return &AdjMatrix{size: size, matrix: m}

}

//Add: create path from i to j with length = l
//	*AdjMatrix.Add(1, 3, 5)
func (a *AdjMatrix) Add(i, j int, l int64) error {
	if a.isOutOfSize(i, j) {
		return ErrSizeLimit
	}
	a.matrix[i][j] = l
	return nil
}

// Length: get length of the path between i and j
// *AdjMatrix.Length(1, 3) = 5
func (a *AdjMatrix) Length(i, j int) int64 {
	return a.matrix[i][j]
}

// Adj: get all vertices that adjacent to i
// *AdjMatrix.Adj(1) = [3]
func (a *AdjMatrix) Adj(i int) []int {
	var result []int
	for j := 0; j < a.size; j++ {
		if a.matrix[i][j] != 0 {
			result = append(result, j)
		}
	}
	return result
}
