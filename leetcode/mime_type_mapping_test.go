package leetcode

import (
	"fmt"
	"testing"
)

const input = `
util.go:                            text/plain
sort/util.go:                       text/plain
structure/binary_tree.go:           text/plain
structure/binary_tree_test.go:      text/plain
structure/btree.go:                 text/plain
structure/btree_test.go:            text/plain
structure/dynamic_programming.go:   text/plain`

func TestFileMimeTypeMapping(t *testing.T) {
	m := FileMimeTypeMapping(input)
	for k, v := range m {
		fmt.Printf("%s: %s\n", k, v)
	}
}
