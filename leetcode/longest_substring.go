package leetcode

func LengthOfLongestSubstring1(s string) int {
	// a(r, l), b, c, b, d: 1
	// a(r), b(l), c, b, d: 2
	// a(r), b, c(l), b, d: 3
	// a(r), b, c, b(l), d
	// a, b(r), c, b(l), d
	// a, b, c(r), b(l), d: 2
	// a, b, c(r), b, d(l): 3
	var left, right, size int
	histogram := make(map[string]int)
	for right < len(s) {
		c1 := string(s[right])
		histogram[c1]++
		if histogram[c1] > 1 {
			for histogram[string(s[right])] > 1 {
				histogram[string(s[left])]--
				left++
			}
		}
		if size < right-left+1 {
			size = right - left + 1
		}
		right++
	}
	return size
}

func LengthOfLongestSubstring2(s string) int {
	// 時間がかかり過ぎる
	var m int
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			if !contain(s, i, j) {
				if j-i > m {
					m = j - i
				}
			}
		}
	}
	return m
}

func contain(s string, i, j int) bool {
	histogram := make(map[string]int)
	substr := s[i:j]
	for _, c := range substr {
		if _, ok := histogram[string(c)]; ok {
			return true
		}
		histogram[string(c)]++
	}
	return false
}

func LengthOfLongestSubstring3(s string) int {
	// 時間がかかり過ぎる
	if len(s) < 2 {
		return len(s)
	}
	var n1, n2, n3, n4 int
	for i, c1 := range s {
		for j, c2 := range s {
			if i < j && c1 == c2 {
				var m1 int
				n1 = LengthOfLongestSubstring3(s[:j])
				n2 = LengthOfLongestSubstring3(s[j:])
				if n1 > n2 {
					m1 = n1
				} else {
					m1 = n2
				}
				var m2 int
				n3 = LengthOfLongestSubstring3(s[:i+1])
				n4 = LengthOfLongestSubstring3(s[i+1:])
				if n3 > n4 {
					m2 = n3
				} else {
					m2 = n4
				}
				if m1 > m2 {
					return m1
				} else {
					return m2
				}
			}
		}
	}
	return len(s)
}
