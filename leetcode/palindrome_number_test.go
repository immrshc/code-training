package leetcode

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		num      int
		expected bool
	}{
		{num: 12321, expected: true},
		{num: 121, expected: true},
		{num: 1, expected: true},
		{num: -121, expected: false},
		{num: 10, expected: false},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%d", test.num), func(t *testing.T) {
			if IsPalindrome1(test.num) != test.expected {
				t.Fatal()
			}
		})
	}
}
