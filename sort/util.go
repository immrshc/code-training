package sort

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

func CheckSort(t *testing.T, f func([]int) []int) {
	inputs := [][]int{
		rand.Perm(15),
		rand.Perm(30),
		rand.Perm(50),
	}
	for i, input := range inputs {
		t.Run(fmt.Sprintf("pattern: %d", i), func(t *testing.T) {
			fmt.Printf("before: %v\n", input)
			output := f(input)
			fmt.Printf("after : %v\n", output)

			sorted := sort.SliceIsSorted(output, func(i, j int) bool {
				return output[i] < output[j]
			})
			if !sorted {
				t.Error("not sorted")
			}
		})
	}
}
