package main

import "fmt"

// https://leetcode.cn/problems/first-unique-character-in-a-string/

func firstUniqChar(s string) int {
	m := make(map[byte]bool)
	for i := 0; i < len(s); i++ {
		k := s[i]
		if _, ok := m[k]; ok {
			m[k] = true
		} else {
			m[k] = false
		}
	}
	for i := 0; i < len(s); i++ {
		if v, ok := m[s[i]]; ok && !v {
			return i
		}
	}
	return -1
}

func main() {
	// 0
	fmt.Println(firstUniqChar("leetcode"))
	// -1
	fmt.Println(firstUniqChar("aaadadaa"))
}
