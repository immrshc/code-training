package sort

import "testing"

func TestHeapSort(t *testing.T) {
	CheckSort(t, HeapSort)
}
