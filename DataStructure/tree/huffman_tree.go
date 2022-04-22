package tree

import (
	"errors"
)

/*
	哈夫曼树主要用于对文本进行编码,压缩文本大小
	哈夫曼树的原则:
		1.权值越大的叶节点越靠近根节点
		2.权值越小的叶节点越远离根节点
		3.只有叶节点用来存数据
		4.没有度为1的节点
		5.n个叶子结点的哈夫曼树共有2n-1个节点
	哈夫曼树用于编码的时候左子树代表0右子树代表1,从根节点走到某叶节点时的路径(0/1)就是该叶节点的编码
*/

// BuildHuffmanTree 从一个slice生成哈夫曼树(数据应按频率的升序存放,频率低的数据放在前面)
func BuildHuffmanTree(s []ElementType) *Node {
	if len(s) <= 1 {
		panic(errors.New("slice need more element"))
	}
	// 建哈夫曼树(s是按使用频率升序排好的数据)
	var head *Node
	for _, v := range s {
		if head == nil {
			// 初次进入循环
			head = createDataNode(v)
		} else {
			head = createConnectNode(head, createDataNode(v))
		}
	}
	return head
}

// createConnectNode 创建一个连接/非叶节点(data为空,度一定为2)
func createConnectNode(left, right *Node) *Node {
	n := new(Node)
	n.Left = left
	n.Right = right
	return n
}

// createDataNode 创建一个数据节点(度一定为0)
func createDataNode(data ElementType) *Node {
	return CreateNode(data, nil, nil)
}
