package sort

func MergeSort(input []int) []int {
	if len(input) < 2 {
		return input
	}
	mid := len(input) / 2
	former := MergeSort(input[:mid])
	latter := MergeSort(input[mid:])

	output := make([]int, 0)
	for {
		if len(former) == 0 {
			output = append(output, latter...)
			break
		}
		if len(latter) == 0 {
			output = append(output, former...)
			break
		}
		if former[0] < latter[0] {
			output = append(output, former[0])
			former = former[1:]
		} else {
			output = append(output, latter[0])
			latter = latter[1:]
		}
	}
	return output
}
