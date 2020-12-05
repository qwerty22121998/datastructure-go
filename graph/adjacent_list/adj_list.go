package adjacent_list

import (
	"errors"
	"github.com/qwerty22121998/datastructure-go/common"
)

// #TODO : write test
type AdjList struct {
	size int
	edge [][]common.PairIntInt64
}

var ErrSizeLimit = errors.New("out of size")

func (a *AdjList) isOutOfSize(i int) bool {
	return i < 0 || i >= a.size
}

func NewAdjList(size int) *AdjList {
	m := make([][]common.PairIntInt64, size)
	return &AdjList{size: size, edge: m}
}

func (a AdjList) Add(i, j int, l int64) error {
	if a.isOutOfSize(i) || a.isOutOfSize(j) {
		return ErrSizeLimit
	}
	a.edge[i] = append(a.edge[i], common.PairIntInt64{First: j, Second: l})
	return nil
}

func (a AdjList) Length(i, j int) int64 {
	if a.isOutOfSize(i) || a.isOutOfSize(j) {
		return -1
	}
	for _, v := range a.edge[i] {
		if v.First == j {
			return v.Second
		}
	}
	return 0
}

func (a AdjList) Adj(i int) []int {
	if a.isOutOfSize(i) {
		return nil
	}
	var result []int
	for _, v := range a.edge[i] {
		result = append(result, v.First)
	}
	return result
}
