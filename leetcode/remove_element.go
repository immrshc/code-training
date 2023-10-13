package leetcode

func RemoveElement(nums []int, val int) int {
	// 指定された値と一致しない要素数を返す
	// 元のスライスの最初のその要素数分は指定された値と一致しない
	// [0, 1, 2, 1, 2] => [0, 1, 1, 1, 2]
	var pos int
	for _, num := range nums {
		if num != val {
			nums[pos] = num
			pos++
		}
	}
	return pos
}
