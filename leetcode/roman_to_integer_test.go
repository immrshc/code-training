package leetcode

import "testing"

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		str      string
		expected int
	}{
		{str: "I", expected: 1},
		{str: "III", expected: 3},
		{str: "IV", expected: 4},
		{str: "VI", expected: 6},
		{str: "XXVII", expected: 27},
		{str: "LVIII", expected: 58},
		{str: "MCMXCIV", expected: 1994},
	}
	for _, test := range tests {
		t.Run(test.str, func(t *testing.T) {
			if n := RomanToInt1(test.str); n != test.expected {
				t.Fatal(n)
			}
		})
	}
}
