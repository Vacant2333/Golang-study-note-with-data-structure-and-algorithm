package main

import "sort"

func main() {
}

func eraseOverlapIntervals(intervals [][]int) int {
	L := len(intervals)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	ans, right := 1, intervals[0][1]
	for i := 0; i < L; i++ {
		if right <= intervals[i][0] {
			ans++
			right = intervals[i][0]
		}
	}
	return L - ans
}
