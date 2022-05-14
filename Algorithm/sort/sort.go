package sort

/*
	冒泡和插入排序中有一个概念叫逆序对,如果某两个元素a[i] > [aj]
	则称i,j是一个逆序对,冒泡和插入排序对同一组数据排序时,效率实际
	上是一样的,因为他们的逆序对的数量是一样的,由此可知,在一组基本
	有序的数据中,冒泡和插入排序是非常高效的(逆序对的数量<n),时间
	复杂度基本在O(n)左右

	定理:任意n个不同元素组成的序列平均具有n(n-1)/4个逆序对,也就是数量级在n^2
	延伸:任何仅以交换相邻两元素的排序算法,其平均时间复杂度为Ω(n),要提高算法效率
		的话,我们每次就要消去更多的逆序对(冒泡和排序每次都只能消去一个)

	排序类型: 升序
	搜索:
		二分查找 done
	排序:
		冒泡排序 done
		插入排序 done
		希尔排序 done
		快速排序
	TODO: 快排
*/

// IsSorted 判断是否已排序(升序)好
func IsSorted(s []int) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] > s[i+1] == true {
			return false
		}
	}
	return true
}

// BubbleSort 冒泡排序(依次扫描相邻的两个元素,小的放左边)
// 时间效率:最好O(n),最坏O(n^2)
// http://c.biancheng.net/view/6506.html
func BubbleSort(s []int) {
	for right := 1; right < len(s); right++ {
		// 如果一趟排序中flag没有变为true就表示数据已排序好,直接break
		flag := false
		// 右边界每次都会-1,因为每一趟排序之后,最大的元素就会被放到最后
		for left := 0; left < len(s)-right; left++ {
			if s[left] > s[left+1] {
				s[left], s[left+1] = s[left+1], s[left]
				flag = true
			}
		}
		if flag == false {
			break
		}
	}
}

// InsertionSort 插入排序(遍历下标从1开始的元素,插入到前面已经排序好的数据中)
// 时间效率:最好O(n),最坏O(n^2)
// https://www.runoob.com/data-structures/insertion-sort.html
func InsertionSort(s []int) {
	for i := 1; i < len(s); i++ {
		// 保存当前元素的值
		tmp := s[i]
		k := i
		// 向左找到一个合适的位置插入当前元素s[i]
		for ; k > 0 && s[k-1] > tmp; k-- {
			// 把已排序好的元素中比s[i]大的部分右移
			s[k] = s[k-1]
		}
		// 插入s[i],下一个是s[i+1]
		s[k] = tmp
	}
}

// ShellSort 希尔排序
func ShellSort(s []int) {
	/*
		基本的希尔排序的增量序列是n,n/2,n/4...
		但是它的时间效率也是O(n^2)
		这里使用Sedgewick增量序列,效率更高
		他的时间效率是O(n^6/5),但是目前无法证明
		这里只写了Sedgewick的一部分增量
	*/
	sedgewick := []int{929, 505, 209, 109, 41, 19, 5, 1, 0}
	sedgewickIndex := 0
	// 增量不能大于数组长度
	for sedgewick[sedgewickIndex] > len(s) {
		sedgewickIndex++
	}
	for d := sedgewick[sedgewickIndex]; d > 0; d = sedgewick[sedgewickIndex] {
		// 插入排序
		for i := d; i < len(s); i++ {
			// 保存当前元素的值
			tmp := s[i]
			k := i
			// 向左找到一个合适的位置插入当前元素s[i]
			for ; k >= d && s[k-1] > tmp; k-- {
				// 把已排序好的元素中比s[i]大的部分右移
				s[k] = s[k-1]
			}
			// 插入s[i],下一个是s[i+1]
			s[k] = tmp
		}
		// 增量取下一个
		sedgewickIndex++
	}
}
