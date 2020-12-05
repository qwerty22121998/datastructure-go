package disjoint_set

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func prepare() []*Node {
	n := []*Node{
		{0, 0, nil, nil, nil},
		{1, 1, nil, nil, nil},
		{2, 2, nil, nil, nil},
		{3, 3, nil, nil, nil},
		{4, 4, nil, nil, nil},
	}
	return n

}

func TestGetRoot(t *testing.T) {
	n := prepare()
	assert.Equal(t, n[0].Root(), n[0])
	assert.Equal(t, n[1].Root(), n[1])
	assert.Equal(t, n[2].Root(), n[2])
	assert.Equal(t, n[3].Root(), n[3])
	assert.Equal(t, n[4].Root(), n[4])
}

func TestJoin(t *testing.T) {
	n := prepare()
	n[0].Join(n[1])
	n[1].Join(n[2])
	assert.Equal(t, n[0].Root(), n[2].Root())
	n[3].Join(n[4])
	assert.Equal(t, n[3].Root(), n[4].Root())
	n[4].Join(n[2])
	assert.Equal(t, n[4].Root(), n[2].Root())
	assert.Equal(t, n[3].Root(), n[2].Root())
}
