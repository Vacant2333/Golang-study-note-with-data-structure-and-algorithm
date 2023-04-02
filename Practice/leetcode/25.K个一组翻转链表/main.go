package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	reverse(dummy, k)
	return dummy.Next
}

func reverse(pre *ListNode, k int) {
	first := pre.Next
	last := pre.Next
	for i := 0; i < k; i++ {
		if last == nil {
			return
		}
		last = last.Next
	}
	cur := pre.Next
	next := cur.Next
	for i := 0; i < k-1; i++ {
		nextTmp := next.Next
		next.Next = cur
		cur = next
		next = nextTmp
	}
	pre.Next = cur
	first.Next = last
	reverse(first, k)
}
