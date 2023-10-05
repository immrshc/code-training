package sort

func HeapSort(input []int) []int {
	// build a heap
	// - add an element into the last index
	// - raise the element up to the relevant position
	var heap []int
	for _, n := range input {
		heap = push(heap, n)
	}
	// fetch the 1st element from the heap
	// - get the 1st element
	// - put the last element on the 1st one
	// - reconstruct the heap
	output := make([]int, 0, len(heap))
	for len(heap) > 0 {
		output = append(output, heap[0])
		heap = pop(heap)
	}
	return output
}

func push(input []int, elm int) []int {
	input = append(input, elm)
	for i := len(input) - 1; i > 0; i = parent(i) {
		if input[parent(i)] <= input[i] {
			break
		}
		input[parent(i)], input[i] = input[i], input[parent(i)]
	}
	return input
}

func pop(input []int) []int {
	input[0] = input[len(input)-1]
	input = input[:len(input)-1]
	for i := 0; i < len(input); {
		if left(i) >= len(input) {
			break
		}
		j := left(i)
		if len(input) > right(i) && input[left(i)] > input[right(i)] {
			j = right(i)
		}
		if input[i] < input[j] {
			break
		}
		input[i], input[j] = input[j], input[i]
		i = j
	}
	return input
}

func parent(idx int) int {
	return (idx - 1) / 2
}

func left(idx int) int {
	return idx*2 + 1
}

func right(idx int) int {
	return left(idx) + 1
}
