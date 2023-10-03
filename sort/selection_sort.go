package sort

func SelectionSort(input []int) []int {
	n := len(input)
	for i := 0; i < n-1; i++ {
		lowkey := i
		lowest := input[i]
		for j := i + 1; j < n; j++ {
			if input[j] < lowest {
				lowkey = j
				lowest = input[j]
			}
		}
		input[i], input[lowkey] = input[lowkey], input[i]
	}
	return input
}
