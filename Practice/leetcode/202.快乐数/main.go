package main

import (
	"strconv"
)

func isHappy(n int) bool {
	m := make(map[int]bool, 10)
	for {
		n = to(n)
		if _, ok := m[n]; ok {
			// 存在这个key 退出
			return false
		} else {
			// 不存在这个key
			m[n] = true
			if n == 1 {
				return true
			}
		}
	}
}

func to(n int) int {
	s := strconv.Itoa(n)
	sum := 0
	for _, v := range s {
		tmp, _ := strconv.Atoi(string(v))
		sum += tmp * tmp
	}
	return sum
}
