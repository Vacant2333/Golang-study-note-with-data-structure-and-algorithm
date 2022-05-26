package segment_tree

/*
	线段树 Segment Tree
	线段树是一棵平衡二叉树。母结点代表整个区间的和，越往下区间越小。
	注意，线段树的每个节点都对应一条线段（区间）
	但并不保证所有的线段（区间）都是线段树的节点，这两者应当区分开。
	两个子节点分别储存 [left, mid] 和 [mid + 1, right] 的和
	不仅可以存和也可以存最大/最小值等数据
	左儿子下标:index*2+1

	https://zhuanlan.zhihu.com/p/106118909
	https://oi-wiki.org/ds/seg/
*/

type segmentTree struct {
	treeSum []int
	// 存区间和,切片的大小一般是元素数量*4,防止溢出
	mark []int
	// 懒/延迟标记(永久化,不会向下传递)
}

func BuildSegmentTree(s []int) segmentTree {
	st := segmentTree{}
	st.treeSum = make([]int, len(s)*4)
	st.mark = make([]int, len(s)*4)
	// 递归处理

	return st
}
