package adjacent_map

type AdjMap struct {
	edge map[int]map[int]int64
}

func NewAdjMap() *AdjMap {
	return &AdjMap{edge: make(map[int]map[int]int64)}
}

func (a *AdjMap) Add(i, j int, l int64) error {
	mp, ok := a.edge[i]
	if !ok {
		mp = make(map[int]int64)
		a.edge[i] = mp
	}
	mp[j] = l
	return nil
}

func (a *AdjMap) Length(i, j int) int64 {
	return a.edge[i][j]
}

func (a *AdjMap) Adj(i int) []int {
	var result []int
	for k := range a.edge[i] {
		result = append(result, k)
	}
	return result
}
