package leetcode

import (
	"strings"
)

func FileMimeTypeMapping(input string) map[string]string {
	mapping := make(map[string]string)
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			continue
		}
		mapping[parts[0]] = strings.TrimSpace(parts[1])
	}
	return mapping
}
