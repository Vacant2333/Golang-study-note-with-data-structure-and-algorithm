package main

// https://leetcode.cn/problems/merge-two-sorted-lists/

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}
	head := &ListNode{}
	cur := head
	for list1 != nil || list2 != nil {
		if list1 != nil && list2 != nil {
			if list1.Val < list2.Val {
				cur.Next = list1
				list1 = list1.Next
			} else {
				cur.Next = list2
				list2 = list2.Next
			}
			cur = cur.Next
		} else if list1 != nil {
			// 接list1
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
			// 接list2
		}
	}
	return head.Next
}

func main() {

}
