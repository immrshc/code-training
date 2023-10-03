package sort

func BubbleSort(input []int) []int {
	n := len(input)
	for i := 0; i < n; i++ {
		for j := n - 1; i < j; j-- {
			if input[j] < input[j-1] {
				input[j], input[j-1] = input[j-1], input[j]
			}
		}
	}
	return input
}
