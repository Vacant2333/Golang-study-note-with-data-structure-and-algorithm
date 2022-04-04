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
	// 根节点
	t := CreateNode(s[0], nil, nil)
	// 插入其余元素
	for i := 1; i < len(s); i++ {
		t.InsertBSTNode(s[i])
	}
	return t
}

// FindBSTNode 搜索BST节点
func (node *Node) FindBSTNode(data ElementType) *Node {
	cur := node
	for cur != nil {
		if data < cur.Data {
			// data比当前节点小,向左走
			cur = cur.Left
		} else if data > cur.Data {
			// data比当前节点大,向右走
			cur = cur.Right
		} else {
			// data和当前节点相等 直接返回
			return cur
		}
	}
	return nil
}

// FindMinBSTNode

// InsertBSTNode BST插入节点(比当前节点小就走左边,否则右边,相等则无操作)
func (node *Node) InsertBSTNode(data ElementType) {
	cur := node
	for cur.Data != data {
		if data < cur.Data {
			// 要插入的节点小于当前节点
			if cur.Left != nil {
				// 左节点不为空 移动t
				cur = cur.Left
			} else {
				// 左节点是空的 插入新节点后退出
				cur.Left = CreateNode(data, nil, nil)
				return
			}
		} else {
			// 要插入的节点大于当前节点
			if cur.Right != nil {
				// 右节点不为空 移动t
				cur = cur.Right
			} else {
				// 左节点是空的 插入新节点后退出
				cur.Right = CreateNode(data, nil, nil)
				return
			}
		}
	}
	// 到了这里就是t和data相等 不作任何操作 直接退出
}
