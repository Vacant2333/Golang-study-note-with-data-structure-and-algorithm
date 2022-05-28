package main

// https://leetcode.cn/problems/valid-anagram/

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sMap := [26]int{}
	tMap := [26]int{}
	for _, c := range s {
		sMap[c-'a']++
	}
	for _, c := range t {
		tMap[c-'a']++
	}
	return sMap == tMap
}

func main() {

}
