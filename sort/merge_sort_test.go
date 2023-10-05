package sort

import "testing"

func TestMergeSort(t *testing.T) {
	CheckSort(t, MergeSort)
}
