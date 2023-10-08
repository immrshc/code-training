package leetcode

import (
	"testing"
)

func TestValidParentheses(t *testing.T) {
	tests := []struct {
		str      string
		expected bool
	}{
		{
			str:      "(",
			expected: false,
		},
		{
			str:      "()",
			expected: true,
		},
		{
			str:      "()[]{}",
			expected: true,
		},
		{
			str:      "(]",
			expected: false,
		},
		{
			str:      "([{}])",
			expected: true,
		},
		{
			str:      "([{]})",
			expected: false,
		},
		{
			str:      "([{}]))",
			expected: false,
		},
		{
			str:      ")([]",
			expected: false,
		},
	}
	for _, test := range tests {
		t.Run(test.str, func(t *testing.T) {
			if ValidParentheses(test.str) != test.expected {
				t.Fatal()
			}
		})
	}
}
