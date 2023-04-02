package Template

import (
	"strconv"
)

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

// strings.Repeat()

func s2i(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
