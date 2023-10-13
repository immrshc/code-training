package leetcode

import (
	"fmt"
	"slices"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		val      int
		nums     []int
		expected []int
	}{
		{
			val:      3,
			nums:     []int{3, 2, 2, 3},
			expected: []int{2, 2},
		},
		{
			val:      2,
			nums:     []int{0, 1, 2, 2, 3, 0, 4, 2},
			expected: []int{0, 0, 1, 3, 4},
		},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("val: %d", test.val), func(t *testing.T) {
			k := RemoveElement(test.nums, test.val)
			if k != len(test.expected) {
				t.Fatalf("expected length (%v) is different from %v", test.expected, k)
			}
			slices.Sort(test.nums[:k])
			for i := 0; i < len(test.expected); i++ {
				if test.nums[i] != test.expected[i] {
					t.Fatalf("expected nums (%v) is different from %v", test.expected, test.nums)
				}
			}
		})
	}
}
