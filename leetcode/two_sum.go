package leetcode

func twoSum1(nums []int, target int) []int {
	// 値から見つけるべき値が分かる
	// そして、その値に対応するインデックスが分かれば良い
	// したがって、値とインデックスのハッシュテーブルを用いるとO(1)で対応するインデックスが分かる
	// しかし、ハッシュテーブルは衝突を考慮しなければならない
	// つまり、異なるインデックスで同じ値が渡された場合を想定する必要がある
	numIdx := make(map[int][]int)
	for i, n := range nums {
		numIdx[n] = append(numIdx[n], i)
	}
	for n, i := range numIdx {
		m := target - n
		if m == n {
			if len(numIdx[m]) == 2 {
				return numIdx[m]
			}
			continue
		}
		if j, ok := numIdx[m]; ok {
			var indices []int
			indices = append(indices, i...)
			indices = append(indices, j...)
			return indices
		}
	}
	return nil
}

func twoSum2(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i == j {
				break
			}
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
