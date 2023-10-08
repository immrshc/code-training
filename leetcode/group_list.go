package leetcode

import (
	"sort"
)

func GroupList(pairs [][]int) [][]int {
	// input: [[1, 2], [4, 8], [7, 11], [3, 7], [5, 12], [2, 4]]
	// output: [[1, 2, 4, 8], [3, 7, 11], [5, 12]]
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0] < pairs[j][0]
	})
	versions := make(map[int][]int)
	for _, pair := range pairs {
		prev, latest := pair[0], pair[1]
		list, ok := versions[prev]
		if ok {
			versions[latest] = append(list, latest)
			delete(versions, prev)
		} else {
			versions[latest] = pair
		}
	}
	outputs := make([][]int, 0)
	for _, v := range versions {
		outputs = append(outputs, v)
	}
	return outputs
}
