package sort

import "testing"

func TestSelectionSort(t *testing.T) {
	CheckSort(t, SelectionSort)
}
