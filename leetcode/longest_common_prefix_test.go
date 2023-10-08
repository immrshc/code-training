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
			strs:     []string{"flower", "flow", "flight"},
			expected: "fl",
		},
		{
			strs:     []string{"flow", "flowed", "flowing"},
			expected: "flow",
		},
		{
			strs:     []string{"a", "b", "c"},
			expected: "",
		},
	}
	for _, test := range tests {
		t.Run(strings.Join(test.strs, ","), func(t *testing.T) {
			if prefix := LongestCommonPrefix2(test.strs); prefix != test.expected {
				t.Fatal(prefix)
			}
		})
	}
}
