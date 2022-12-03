package sort

// BinarySearch 二分查找(s必须是排序好的(升序)) 返回下标
func BinarySearch(s []int, value int) int {
	left, right := 0, len(s)
	for left < right {
		mid := (left + right) / 2
		if s[mid] > value {
			// value在mid的左边,将right向左移动,到mid-1
			right = mid - 1
		} else if s[mid] < value {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

// BinarySearch2 更好理解的二分查找 https://www.bilibili.com/video/BV1d54y1q7k7
func BinarySearch2(s []int, target int) int {
	// left,right初始时在slice的外面
	left, right := -1, len(s)
	for left+1 != right {
		// 当left和right贴在一起(也就是他们中间没有空位)了就返回
		mid := (left + right) / 2
		if s[mid] <= target {
			left = mid
		} else {
			right = mid
		}
	}
	/*
		这种写法是寻找第一个小于等于target的目标,左右区间要分清楚
		如果想寻找第一个>=target的目标(也就是left的元素都是小于target的值), if s[mid] < target, return right
		如果想寻找最后一个<target的目标(也就是left的元素都是小于target的值), if s[mid] < target, return left
	*/
	return left
}
