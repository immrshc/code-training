package leetcode

import (
	"strings"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		strs     []string
		expected string
	}{
		{
			strs:     []string{},
			expected: "",
		},
	}
	for _, test := range tests {
		t.Run(strings.Join(test.strs, ","), func(t *testing.T) {
			if prefix := LongestCommonPrefix(test.strs); prefix != test.expected {
				t.Fatal(prefix)
			}
		})
	}
}
