package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

func TestGroupList(t *testing.T) {
	tests := []struct {
		input  [][]int
		output [][]int
	}{
		{
			input:  [][]int{{1, 2}, {4, 8}, {7, 11}, {3, 7}, {5, 12}, {2, 4}},
			output: [][]int{{1, 2, 4, 8}, {3, 7, 11}, {5, 12}},
		},
		{
			input:  [][]int{{1, 2}, {3, 4}},
			output: [][]int{{1, 2}, {3, 4}},
		},
		{
			input:  [][]int{{1, 2}, {2, 4}},
			output: [][]int{{1, 2}, {2, 4}},
		},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("input: %v", test.input), func(t *testing.T) {
			result := GroupList(test.input)
			sort.Slice(result, func(i, j int) bool {
				return result[i][0] < result[j][0]
			})
			if fmt.Sprintf("%v", result) != fmt.Sprintf("%v", test.output) {
				t.Fatal(result)
			}
		})
	}
}
