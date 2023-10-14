package leetcode

import "testing"

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		str      string
		expected string
	}{
		{str: "a", expected: "a"},
		{str: "ab", expected: "a"},
		{str: "abcba", expected: "abcba"},
		{str: "abacab", expected: "bacab"},
		{str: "aabcdedcba", expected: "abcdedcba"},
		{str: "abba", expected: "abba"},
		{str: "bbc", expected: "bb"},
		{str: "bbbc", expected: "bbb"},
		{str: "bbbbc", expected: "bbbb"},
	}
	for _, test := range tests {
		t.Run(test.str, func(t *testing.T) {
			if r := LongestPalindrome2(test.str); r != test.expected {
				t.Fatalf("result: %s, but expected: %s", r, test.expected)
			}
		})
	}
}
