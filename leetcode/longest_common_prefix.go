package leetcode

import (
	"strings"
)

func LongestCommonPrefix1(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := ""
	str := strs[0]
	for _, char := range str {
		next := prefix + string(char)
		for _, s := range strs[1:] {
			if !strings.HasPrefix(s, next) {
				return prefix
			}
		}
		prefix = next
	}
	return prefix
}

func LongestCommonPrefix2(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	middle := len(strs) / 2
	prefix1 := LongestCommonPrefix2(strs[:middle])
	prefix2 := LongestCommonPrefix2(strs[middle:])
	var merged string
	for _, runes := range prefix1 {
		next := merged + string(runes)
		if !strings.HasPrefix(prefix2, next) {
			return merged
		}
		merged = next
	}
	return merged
}
