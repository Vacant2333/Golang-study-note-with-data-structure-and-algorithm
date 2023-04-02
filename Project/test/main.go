package main

import "fmt"

func main() {
	var count int
	fmt.Scan(&count)
	boxs := map[byte][]int{}
	for i := 0; i < count; i++ {
		var char byte
		fmt.Scanf("%c", &char)
		if _, ok := boxs[char]; ok {
			// 如果这个字符存过
			boxs[char] = append(boxs[char], i)
		} else {
			boxs[char] = []int{i}
		}
	}
	dp := map[byte][]int{}
	for startChar := range boxs {

	}

}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
