package main

// https://leetcode.cn/problems/remove-linked-list-elements/

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	// 烧饼节点
	pre := &ListNode{0, head}
	cur := pre.Next
	preCur := pre

	for cur != nil {
		if cur.Val == val {
			// 删除cur
			preCur.Next = cur.Next
			cur = preCur.Next
		} else {
			preCur = cur
			cur = cur.Next
		}
	}
	return pre.Next
}

func main() {

}
