package main

import "sort"

func largestWordCount(messages []string, senders []string) string {
	m := make(map[string]int)
	n := make([]string, 0)
	for i := 0; i < len(senders); i++ {
		m[senders[i]] = 0
		n = append(n, senders[i])
	}
	sort.Strings(n)
	for i := 0; i < len(messages); i++ {
		m[senders[i]] += c(messages[i])
	}
	maxIndex := n[0]
	maxIndexs := make([]string, 0)
	for name, count := range m {
		if count > m[maxIndex] {
			maxIndex = name
			maxIndexs = []string{name}
		}
		if count == m[maxIndex] {
			maxIndexs = append(maxIndexs, name)
		}
	}
	for i := len(n) - 1; i >= 0; i-- {
		if in(maxIndexs, n[i]) {
			return n[i]
		}
	}
	return ""
}

func in(a []string, b string) bool {
	for i := 0; i < len(a); i++ {
		if b == a[i] {
			return true
		}
	}
	return false
}

func c(msg string) int {
	co := 0
	for _, v := range msg {
		if v == ' ' {
			co++
		}
	}
	return co + 1
}

func main() {

}
