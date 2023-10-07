package leetcode

import (
	"strconv"
)

func IsPalindrome1(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	if x%10 == 0 {
		return false
	}
	// 12321の場合のpartの遷移
	// 0 + 1 = 1
	// 10 + 2 = 12
	// 120 + 3 = 123
	// 1230 + 1 = 1231
	var reverse int
	for x > reverse {
		reverse = reverse*10 + x%10
		x /= 10
	}
	return x == reverse || x == reverse/10
}

func IsPalindrome2(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	if x%10 == 0 {
		return false
	}
	str := strconv.Itoa(x)
	for i, j := 0, len(str)-1; i <= len(str)/2; i, j = i+1, j-1 {
		if str[i] != str[j] {
			return false
		}
	}
	return true
}
