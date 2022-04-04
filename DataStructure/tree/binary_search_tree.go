package tree

/*
	二叉搜索树(BST)的实现,基于二叉树(binary_tree)
	这种数据结构本身很少使用,因为他的各种操作复杂度是O(层数),只有在它大致平衡(所有椰子的深度趋于相同)
	才具有优秀的复杂度(当它是完全二叉树时,为O(log n)),然而这只是理想情况
	假如我们依次加入6,5,4,3,2,1 BST的就会退化成链表,复杂度变为O(n)
*/

// CreateBSTFromSlice 从切片生成一个搜索二叉树BST
func CreateBSTFromSlice(s []ElementType) *Node {
	if len(s) == 0 {
		return nil
	}
	t := CreateNode(s[0], nil, nil)
	// 插入其余元素

	return t
}

// InsertBSTNode BST插入节点(比当前节点小就走左边,否则右边,相等则无操作)
func InsertBSTNode(t *Node, data ElementType) {

}
