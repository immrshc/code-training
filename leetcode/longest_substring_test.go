package leetcode

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		str    string
		length int
	}{
		{str: "abbaca", length: 3},
		{str: "dvdf", length: 3},
		{str: "abcb", length: 3},
	}
	for _, test := range tests {
		t.Run(test.str, func(t *testing.T) {
			if result := LengthOfLongestSubstring1(test.str); result != test.length {
				t.Fatalf("result: %d, expected: %d", result, test.length)
			}
		})
	}
}
