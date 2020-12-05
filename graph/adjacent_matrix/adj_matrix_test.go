package adjacent_matrix

import (
	"github.com/qwerty22121998/datastructure-go/graph"
	"github.com/stretchr/testify/assert"
	"testing"
)

func prepareNewMatrix() graph.IGraphInt {
	return NewAdjMatrix(5)
}

func prepareAddEdge(m graph.IGraphInt) {
	if err := m.Add(1, 2, 1); err != nil {
		panic(err)
	}
	if err := m.Add(1, 4, 3); err != nil {
		panic(err)
	}
	if err := m.Add(4, 3, 5); err != nil {
		panic(err)
	}
}

func TestAdjMatrix_Add(t *testing.T) {
	m := prepareNewMatrix()
	if err := m.Add(1, 2, 1); err != nil {
		t.Error(err)
	}
	if err := m.Add(1, 4, 1); err != nil {
		t.Error(err)
	}
	if err := m.Add(4, 3, 1); err != nil {
		t.Error(err)
	}
	expected := [][]int64{
		{0, 0, 0, 0, 0},
		{0, 0, 1, 0, 1},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0},
	}
	assert.Equal(t, expected, m.(*AdjMatrix).matrix)

	err := m.Add(-1, -1, 0)
	assert.Equal(t, ErrSizeLimit, err)
}

func TestAdjMatrix_Adj(t *testing.T) {
	m := prepareNewMatrix()
	prepareAddEdge(m)
	assert.Equal(t, []int{2, 4}, m.Adj(1))
	assert.Equal(t, []int{3}, m.Adj(4))
}

func TestAdjMatrix_Length(t *testing.T) {
	m := prepareNewMatrix()
	prepareAddEdge(m)
	assert.Equal(t, int64(1), m.Length(1, 2))
	assert.Equal(t, int64(3), m.Length(1, 4))
	assert.Equal(t, int64(5), m.Length(4, 3))
}
