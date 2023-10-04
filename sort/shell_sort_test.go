package sort

import (
	"testing"
)

func TestShellSort(t *testing.T) {
	CheckSort(t, ShellSort)
}
