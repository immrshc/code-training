package sort

import (
	"math"
)

func ShellSort(input []int) []int {
	n := len(input)
	seq := make([]int, 0)
	for i := 0; term(i) < n; i++ {
		seq = append(seq, term(i))
	}
	for i := len(seq); i > 0; i-- {
		h := seq[i-1]
		insertionSort(input, h)
	}
	return input
}

func insertionSort(input []int, h int) {
	for i := h; i < len(input); i = i + h {
		for j := i; j >= h && input[j-h] > input[j]; j = j - h {
			input[j], input[j-h] = input[j-h], input[j]
		}
	}
}

func term(n int) int {
	// 4^n + 3*2^(n-1) + 1
	// 1, 8, 23, 77, 281, ...
	if n < 1 {
		return 1
	}
	return int(math.Pow(4, float64(n))) + 3*int(math.Pow(2, float64(n-1))) + 1

}
