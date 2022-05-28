package main

// https://leetcode.cn/problems/ransom-note/

func canConstruct(ransomNote string, magazine string) bool {
	m := [26]int{}
	for _, c := range magazine {
		m[c-'a']++
	}
	for _, c := range ransomNote {
		if m[c-'a'] == 0 {
			return false
		} else {
			m[c-'a']--
		}
	}
	return true
}

func main() {

}
