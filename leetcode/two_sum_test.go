package leetcode

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
	}{
		{
			nums:   []int{3, 2, 4},
			target: 6,
		},
		{
			nums:   []int{1, 3, 4, 2},
			target: 6,
		},
		{
			nums:   []int{3, 3},
			target: 6,
		},
		{
			nums:   []int{-3, 4, 3, 90},
			target: 0,
		},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("target: %d, nums: %v", test.target, test.nums), func(t *testing.T) {
			output := twoSum1(test.nums, test.target)
			t.Log(output)
			if len(output) != 2 || test.nums[output[0]]+test.nums[output[1]] != test.target {
				t.Fatal()
			}
		})
	}
}
