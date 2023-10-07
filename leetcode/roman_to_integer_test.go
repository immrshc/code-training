package leetcode

import "testing"

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		str      string
		expected int
	}{
		{str: "", expected: 0},
	}
	for _, test := range tests {
		t.Run(test.str, func(t *testing.T) {
			if RomanToInt(test.str) != test.expected {
				t.Fatal()
			}
		})
	}
}
