package leetcode

func LongestPalindrome1(s string) string {
	if len(s) < 2 {
		return s
	}
	var sub string
	for first := 0; first < len(s); first++ {
		for last := len(s); last > first; last-- {
			if check(s, first, last) {
				if len(sub) < len(s[first:last]) {
					sub = s[first:last]
				}
			}
		}
	}
	return sub
}

func check(s string, i, j int) bool {
	left, right := i, j-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func LongestPalindrome2(s string) string {
	// □△○△□: 奇数で最内側は1つ
	// □△○○△□: 偶数で最内側は2つ
	// dp[i][j]: i~jが回文になっているかどうか
	// 両端が同じで内側も回文になっている場合はその両端の区間も回文となる
	// s[i] == s[j] && dp[i+1][j-1] => dp[i][j] = true
	n := len(s)
	dp := make(map[int]map[int]bool)
	var sub string

	// 奇数文字の回文の最小単位をtrueにしておく
	for i := 0; i < n; i++ {
		dp[i] = map[int]bool{i: true}
		if len(sub) < 1 {
			sub = s[i : i+1]
		}
	}

	// 偶数文字の回文の最小単位をtrueにしておく
	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = true
			if len(sub) < 2 {
				sub = s[i : i+2]
			}
		}
	}

	for diff := 2; diff < n; diff++ {
		for i := 0; i+diff < n; i++ {
			j := i + diff
			if s[i] == s[j] && dp[i+1][j-1] {
				dp[i][j] = true
				if len(sub) < j-i+1 {
					sub = s[i : j+1]
				}
			}
		}
	}

	return sub
}

func LongestPalindrome3(s string) string {
	if len(s) < 2 {
		return s
	}
	var sub string
	for i := 1; i < len(s); i++ {
		var j, k int
		for j, k = i, i; j >= 0 && k < len(s) && s[j] == s[k]; j, k = j-1, k+1 {
		}
		j, k = j+1, k-1
		if j != k && len(sub) < k-j+1 {
			sub = s[j : k+1]
		}
		if s[i-1] != s[i] {
			continue
		}
		for j, k = i-2, i+1; j >= 0 && k < len(s) && s[j] == s[k]; j, k = j-1, k+1 {
		}
		j, k = j+1, k-1
		if j != k && len(sub) < k-j+1 {
			sub = s[j : k+1]
		}
	}
	if sub == "" {
		return string(s[0])
	}
	return sub
}
