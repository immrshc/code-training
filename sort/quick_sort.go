package sort

func QuickSort(input []int) []int {
	quickSort(input, 0, len(input)-1)
	return input
}

func quickSort(input []int, l, r int) {
	if l >= r {
		return
	}
	v := partition(input, l, r)
	quickSort(input, l, v-1)
	quickSort(input, v+1, r)
}

func partition(input []int, l, r int) int {
	pivot := input[r]
	i, j := l-1, r
	for {
		// init → condition → increment → condition → increment → ...
		// https://go-tour-jp.appspot.com/flowcontrol/1
		for i++; input[i] < pivot; i++ {
		}
		for j--; i < j && input[j] > pivot; j-- {
		}
		if i >= j {
			break
		}
		input[i], input[j] = input[j], input[i]
	}
	input[i], input[r] = input[r], input[i]
	return i
}
