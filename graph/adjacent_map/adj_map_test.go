package adjacent_map

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func prepare() *AdjMap {
	return NewAdjMap()
}

func TestAdjMap_Add(t *testing.T) {
	mp := prepare()
	if err := mp.Add(1, 5, 10); err != nil {
		t.Error(err)
	}
	if err := mp.Add(1, 7, 20); err != nil {
		t.Error(err)
	}
}

func TestAdjMap_Length(t *testing.T) {
	mp := prepare()
	mp.Add(1, 5, 20)
	mp.Add(2, 4, 10)
	assert.Equal(t, int64(20), mp.Length(1, 5))
}

func TestAdjMap_Adj(t *testing.T) {
	mp := prepare()
	mp.Add(1, 2, 20)
	mp.Add(1, 5, 1)
	mp.Add(1, 7, 2)
	assert.Equal(t, []int{2,5,7}, mp.Adj(1))
}