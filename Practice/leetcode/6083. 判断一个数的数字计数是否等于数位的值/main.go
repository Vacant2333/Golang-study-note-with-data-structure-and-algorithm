package main

import (
	"fmt"
)

// https://leetcode.cn/contest/biweekly-contest-79/problems/check-if-number-has-equal-digit-count-and-digit-value/

func digitCount(num string) bool {
	for i, v := range num {
		if int(v-48) != count(num, int32(i+48)) {
			return false
		}
	}
	return true
}

func count(num string, v2 int32) int {
	c := 0
	for _, v := range num {
		if v == v2 {
			c++
		}
	}
	return c
}

func main() {
	// true
	fmt.Println(digitCount("1210"))
}
