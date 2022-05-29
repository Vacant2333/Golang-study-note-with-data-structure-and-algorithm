package main

import "fmt"

func rearrangeCharacters(s string, target string) int {
	targetMap := [26]int32{}
	for _, c := range target {
		targetMap[c-'a']++
	}
	sMap := [26]int32{}
	for _, c := range s {
		sMap[c-'a']++
	}
	count := 0
	flag := true
	for flag {
		for i := 0; i < 26; i++ {
			sMap[i] -= targetMap[i]
			if sMap[i] < 0 {
				flag = false
				break
			}
		}
		if !flag {
			break
		}
		count++
	}

	return count
}

func main() {
	fmt.Println(rearrangeCharacters("ilovecodingonleetcode", "code"))
}
