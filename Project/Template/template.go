package Template

import (
	"strconv"
)

func min(a ...int) int {
	res := a[0]
	for _, v := range a[1:] {
		if v < res {
			res = v
		}
	}
	return res
}

func max(a ...int) int {
	res := a[0]
	for _, v := range a[1:] {
		if v > res {
			res = v
		}
	}
	return res
}

// strings.Repeat()

func s2i(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
