package main

import "fmt"

/*
https://leetcode.cn/problems/longest-palindromic-substring/

*/

func main() {
	fmt.Println(longestPalindrome("aaaa"))
}

func longestPalindrome(s string) string {
	L, maxLen, start, end := len(s), 0, 0, 0
	dp := make([][]bool, L)
	for h := 0; h < L; h++ {
		dp[h] = make([]bool, L)
		dp[h][h] = true
	}

	for l := L - 1; l >= 0; l-- {
		for r := l + 1; r < L; r++ {
			//fmt.Printf("out: [%d, %d] [%s]\n", l, r, s[l:r+1])
			if r-l+1 == 2 {
				// 2 chars
				if s[r] == s[l] {
					dp[l][r] = true
					if 2 > maxLen {
						start, end, maxLen = l, r, 2
					}
				}
			} else {
				// 2+ chars
				if dp[l+1][r-1] && s[l] == s[r] {
					dp[l][r] = true
					if r-l+1 > maxLen {
						start, end, maxLen = l, r, r-l+1
					}
				}
			}
			//fmt.Printf("end: %v\n", dp[l][r])
		}
	}
	return s[start : end+1]
}
