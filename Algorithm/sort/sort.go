package sort

/*
	排序类型: 升序
	搜索:
		二分查找 done
	排序:
		冒泡排序 done
		插入排序 done
		快速排序
	TODO: 快排
*/

// ElementType 可排序/查找的类型的接口
// 已完成: []int,
type ElementType interface {
	// 获得长度
	len() int
	// 通过两个下标比较大小 s[a] > s[b] 则是true
	compare(a int, b int) bool
	// 交换两个下标对应的值
	swap(a int, b int)
}

// IsSorted 判断是否已排序(升序)好
func IsSorted(s ElementType) bool {
	for i := 0; i < s.len()-1; i++ {
		if s.compare(i, i+1) == true {
			return false
		}
	}
	return true
}

// BubbleSort 冒泡排序(依次扫描相邻的两个元素,小的放左边)
// 时间效率:最好O(n),最坏O(n^2)
// http://c.biancheng.net/view/6506.html
func BubbleSort(s ElementType) {
	for right := 1; right < s.len(); right++ {
		// 如果一趟排序中flag没有变为true就表示数据已排序好,直接break
		flag := false
		// 右边界每次都会-1,因为每一趟排序之后,最大的元素就会被放到最后
		for left := 0; left < s.len()-right; left++ {
			if s.compare(left, left+1) {
				s.swap(left, left+1)
				flag = true
			}
		}
		if flag == false {
			break
		}
	}
}

// InsertionSort  插入排序(遍历下标从1开始的元素,插入到前面已经排序好的数据中)
// 时间效率:最好O(n),最坏O(n^2)
// https://www.runoob.com/data-structures/insertion-sort.html
func InsertionSort(s ElementType) {
	for i := 1; i < s.len(); i++ {
		for k := i; k > 0; k-- {
			if s.compare(k, k-1) == false {
				s.swap(k, k-1)
			}
		}
	}
}
