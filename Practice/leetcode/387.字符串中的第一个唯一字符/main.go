package main

import "fmt"

// https://leetcode.cn/problems/first-unique-character-in-a-string/

func firstUniqChar(s string) int {
	m := [26]int{}
	for i := 0; i < len(s); i++ {
		m[s[i]-'a'] += 1
	}
	for i, c := range s {
		if m[c-'a'] == 1 {
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
