package tree

/*
	二叉搜索树(BST)的实现,基于二叉树(binary_tree)
	性质:
		1.非空左子树上所有节点的值均小于它的根节点的值
		2.非空右子树上所有节点的值均大于它的根节点的值
		3.左右子树也分别为二叉搜索树
	特性:
		1.中序遍历下输出的二叉搜索树的节点的值是有序的(升序)
		2.没有重复元素
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
func (root *Node) FindBSTNode(data ElementType) *Node {
	cur := root
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

// FindMinBSTNode 搜索BST最小节点(也可以传入BST的某个节点,把这个节点当成一个BST,然后找它的最小的节点,最左边的节点就是最小的节点)
func (root *Node) FindMinBSTNode() *Node {
	cur := root
	for cur != nil && cur.Left != nil {
		// 一直向左移动 直到它没有左节点
		cur = cur.Left
	}
	return cur
}

// FindMaxBSTNode 搜索BST最大节点,同上,最右边的节点就是最大的节点
func (root *Node) FindMaxBSTNode() *Node {
	cur := root
	for cur != nil && cur.Right != nil {
		// 一直向右移动 直到它没有右节点
		cur = cur.Right
	}
	return cur
}

// InsertBSTNode BST插入节点(比当前节点小就走左边,否则右边,相等则无操作)
func (root *Node) InsertBSTNode(data ElementType) {
	cur := root
	for cur.Data != data {
		if data < cur.Data {
			// 要插入的节点小于当前节点
			if cur.Left != nil {
				// 左节点不为空 移动cur
				cur = cur.Left
			} else {
				// 左节点是空的 插入新节点后退出
				cur.Left = CreateNode(data, nil, nil)
				return
			}
		} else {
			// 要插入的节点大于当前节点
			if cur.Right != nil {
				// 右节点不为空 移动cur
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

// DeleteBSTNode 删除BST节点
func (root *Node) DeleteBSTNode(data ElementType) {
	var pre, cur *Node
	cur = root
	for cur != nil {
		if data < cur.Data {
			// 移动cur到要删除的节点
			pre = cur
			cur = cur.Left
		} else if data > cur.Data {
			// 移动cur到要删除的节点
			pre = cur
			cur = cur.Right
		} else {
			// cur是要删除的节点
			if cur.CountSon() == 0 {
				// 要删除的节点没有子节点 直接从父节点断开连接
				if pre.Left == cur {
					pre.Left = nil
				} else {
					pre.Right = nil
				}
			} else if cur.CountSon() == 1 {
				// 要删除的节点有一个子节点 直接用子节点代替要删除的节点
				var tmp *Node
				if cur.Left != nil {
					// 要删除的节点有一个右节点
					tmp = cur.Right
				} else {
					// 要删除的节点有一个左节点
					tmp = cur.Right
				}
				// 直接用要删除的节点的子节点替换要删除的节点的内容
				cur.Data = tmp.Data
				cur.Left = tmp.Left
				cur.Right = tmp.Right
			} else {
				// 要删除的节点有两个子节点(该节点的右子树的最小值或左子树的最大值来代替他,然后删除最小/最大值)
				rightMin := cur.Right.FindMinBSTNode()
				// 直接用右子树的最小值替换当前要删除的值
				cur.Data = rightMin.Data
				// 删除右子树中的最小值(也就是重复值)
				cur.Right.DeleteBSTNode(cur.Data)
			}
			return
		}
	}
}

// IsValidBST 检查BST是否正确(中序遍历法)
func (root *Node) IsValidBST() bool {
	// 前一个节点 每个节点都小于前一个节点就是正确的BST
	var pre *Node
	var travel func(n *Node) bool
	travel = func(n *Node) bool {
		if n == nil {
			return true
		} else {
			leftResult := travel(n.Left)
			if pre != nil && pre.Data >= n.Data {
				return false
			}
			pre = n
			rightResult := travel(n.Right)
			return leftResult && rightResult
		}
	}
	return travel(root)
}

// IsEqual 判断两个树是否相等
func (root *Node) IsEqual(root2 *Node) bool {

}
