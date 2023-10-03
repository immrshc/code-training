package sort

func InsertionSort(input []int) []int {
	n := len(input)
	for i := 1; i < n; i++ {
		for j := i; j > 0 && input[j-1] > input[j]; j-- {
			input[j], input[j-1] = input[j-1], input[j]
		}
	}
	return input
}
