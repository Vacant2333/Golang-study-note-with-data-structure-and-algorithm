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

// string 转 int
func s2i(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

// GetSumN 求1+2+3+n
func GetSumN(n int) int {
	if n%2 == 0 {
		return (n / 2) * (n + 1)
	} else {
		return (n/2)*(n+1) + (n/2 + 1)
	}
}
