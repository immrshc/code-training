package leetcode

import "testing"

func TestStrStr(t *testing.T) {
	tests := []struct {
		haystack string
		needle   string
		expected int
	}{
		{
			haystack: "sadbutsad",
			needle:   "sad",
			expected: 0,
		},
		{
			haystack: "leetcode",
			needle:   "leeto",
			expected: -1,
		},
		{
			haystack: "sadbutsad",
			needle:   "a",
			expected: 1,
		},
		{
			haystack: "abcdefg",
			needle:   "g",
			expected: 6,
		},
	}
	for _, test := range tests {
		t.Run(test.haystack, func(t *testing.T) {
			if idx := StrStr(test.haystack, test.needle); idx != test.expected {
				t.Fatalf("%s is matched in %s at %d, but %d", test.needle, test.haystack, test.expected, idx)
			}
		})
	}
}
