package sort

func BinSort(input []int) []int {
	// 要素の長さがnの場合[0,n)の重複の無いスライスの前提
	bin := make([]int, len(input))
	for _, i := range input {
		bin[i] = i
	}
	return bin
}
