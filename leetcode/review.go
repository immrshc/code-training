package leetcode

import "strings"

func twoSum(nums []int, target int) []int {
	// n => target - n = m => mのindex
	mapping := make(map[int][]int)
	for i, n := range nums {
		mapping[n] = append(mapping[n], i)
	}
	for i1, n := range nums {
		m := target - n
		indices, ok := mapping[m]
		if !ok {
			continue
		}
		for _, i2 := range indices {
			if i2 != i1 {
				return []int{i1, i2}
			}
		}
	}
	return nil
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	if x%10 == 0 {
		return false
	}
	// 1221, 12321
	for y := 0; y < x; {
		y = y*10 + x%10
		x = x / 10
		if y == x || y == x/10 {
			return true
		}
	}
	return false
}

func romanToInt(s string) int {
	// 小→大であれば大-小
	// 大→小であれば大+小
	var prev, sum int
	for _, c := range s {
		n := SymbolValue[string(c)]
		if prev != 0 && prev < n {
			sum -= prev
			sum += n - prev
		} else {
			sum += n
		}
		prev = n
	}
	return sum
}

func longestCommonPrefix(strs []string) string {
	// 最大公約数的な問題は分割して段階的にマージできる
	n := len(strs)
	if n == 0 {
		return ""
	}
	if n == 1 {
		return strs[0]
	}
	pivot := n / 2
	p1 := longestCommonPrefix(strs[:pivot])
	p2 := longestCommonPrefix(strs[pivot:])

	var prefix string
	for _, runes := range p1 {
		next := prefix + string(runes)
		if !strings.HasPrefix(p2, next) {
			return prefix
		}
		prefix = next
	}
	return prefix
}

func mergeTwoLists(list1, list2 *ListNode) *ListNode {
	// 先頭を選んで、残りを後ろに繋げるを繰り返す
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}
}

func strStr(haystack, needle string) int {
	h := len(haystack)
	n := len(needle)
	if h < n {
		return -1
	}
	for i := 0; i < h-n+1; i++ {
		if haystack[i:i+n] == needle {
			return i
		}
	}
	return -1
}
