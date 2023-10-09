package leetcode

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		nums    []int
		typeNum int
	}{
		{
			nums:    []int{1},
			typeNum: 1,
		},
		{
			nums:    []int{1, 1},
			typeNum: 1,
		},
		{
			nums:    []int{1, 1, 2},
			typeNum: 2,
		},
		{
			nums:    []int{1, 1, 2, 3, 3, 3},
			typeNum: 3,
		},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("nums: %v", test.nums), func(t *testing.T) {
			if pos := RemoveDuplicates(test.nums); pos != test.typeNum {
				t.Fatal(pos)
			}
		})
	}
}
