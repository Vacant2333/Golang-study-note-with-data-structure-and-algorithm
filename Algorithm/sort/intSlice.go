package sort

/*
	在这个文件中实现sort.go中slice的接口([]int)
*/

type IntSlice []int

// 获得长度
func (s IntSlice) len() int {
	return len(s)
}

// 比较两个element a > b 就返回true
func (s IntSlice) compare(a int, b int) bool {
	return s[a] > s[b]
}

// 交换两个element
func (s IntSlice) swap(a, b int) {
	s[a], s[b] = s[b], s[a]
}
