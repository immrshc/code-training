package leetcode

func RemoveDuplicates(nums []int) int {
	prev := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != prev {
			nums[count] = nums[i]
			count++
		}
		prev = nums[i]
	}
	return count
}
