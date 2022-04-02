package sort

/*
	排序类型: 升序
	搜索:
		二分查找 done
	排序:
		冒泡排序 待完成
		快速排序 待完成
*/

// 可排序/查找的类型的接口
// 已完成: []int,
type data interface {
	// 获得长度
	len() int
	// 通过两个下标比较大小 s[a] > s[b] 则是true
	compare(a int, b int) bool
	// 交换两个下标对应的值
	swap(a int, b int)
}

// IsSorted 判断是否已排序好
func IsSorted(s data) bool {
	for i := 0; i < s.len()-1; i++ {
		if s.compare(i, i+1) == true {
			return false
		}
	}
	return true
}

// BubbleSort 冒泡排序
func BubbleSort(s data) {
	// flag: 在内部循环没有操作时退出外部循环
	flag := true
	for flag {
		flag = false
		for i := 0; i < s.len()-1; i++ {
			if s.compare(i, i+1) {
				s.swap(i, i+1)
				flag = true
			}
		}
	}
}
