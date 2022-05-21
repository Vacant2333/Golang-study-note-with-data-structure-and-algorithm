package sort

import (
	"math/rand"
	"time"
)

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
		堆排序   done
		猴子排序 done
		睡眠排序 done
		归并排序
		快速排序
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
		https://baike.baidu.com/item/%E5%B8%8C%E5%B0%94%E5%A2%9E%E9%87%8F/6853339?fr=aladdin
	*/
	sedgewick := []int{929, 505, 209, 109, 41, 19, 5, 1, 0}
	sedgewickIndex := 0
	// 增量不能大于数组长度
	for sedgewick[sedgewickIndex] > len(s) {
		sedgewickIndex++
	}
	// 可以把插入排序的增量理解为1,这里换成sedgewick的元素,一直向下取直到0
	for distance := sedgewick[sedgewickIndex]; distance > 0; distance = sedgewick[sedgewickIndex] {
		// 插入排序
		for i := distance; i < len(s); i++ {
			tmp := s[i]
			k := i
			for ; k >= distance && s[k-distance] > tmp; k -= distance {
				s[k] = s[k-distance]
			}
			s[k] = tmp
		}
		// 增量取下一个
		sedgewickIndex++
	}
}

// HeapSort 堆排序
// 实际使用没有用sedgewick序列的希尔排序效果好
func HeapSort(s []int) {
	// shiftDown 堆排序,下滤(原版可看DataStructure/heap)
	shiftDown := func(root int, N int) {
		// 将前N个元素以root为根的子堆调整为最大堆
		var parent, child, X int
		X = s[root]
		for parent = root; parent*2+1 < N; parent = child {
			child = parent*2 + 1
			// 如果存在右节点且右节点更大, child指向更大的节点(右节点)
			if child <= N-2 && s[child] < s[child+1] {
				child++
			}
			if X >= s[child] {
				break
			} else {
				s[parent] = s[child]
			}
		}
		s[parent] = X
	}

	var i int

	for i = len(s)/2 - 1; i >= 0; i-- {
		shiftDown(i, len(s))
	}
	for i = len(s) - 1; i > 0; i-- {
		// 把最大值放到最后
		s[0], s[i] = s[i], s[0]
		shiftDown(0, i)
	}
}

// BogoSort 猴子排序 平均复杂度:O(n*n!)(这是我见过效率最低的算法)
// 每次把原始数据打乱一次,如果排序好了就返回,没排序好就继续打乱
func BogoSort(s []int) {
	// 打乱这个数组的元素(洗牌)
	rand.Shuffle(len(s), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})
	// 没排序好,再跑一次 :(
	if IsSorted(s) == false {
		BogoSort(s)
	}
}

// SleepSort 睡眠排序 复杂度:????(效率和上面那个猴子排序有的一拼)
// 原理:先睡醒的线程他所表示的元素就是比较小的那个,直接传回去
// 不能有负数,如果有负数可以通过加偏移量来处理!!
// 如果sleep设置过小可能会有误差导致排序失败
func SleepSort(s []int) {
	// 协程传回元素(越小的元素sleep越短,位置就越靠前)
	outChan := make(chan int, len(s))
	// 睡眠协程,传入元素和chan
	sleepRoutine := func(x int) {
		// 等待时间: 元素大小 * 500ms
		time.Sleep(time.Millisecond * 500 * time.Duration(x))
		outChan <- x
	}
	// 启动协程,进行排序
	for i := 0; i < len(s); i++ {
		go sleepRoutine(s[i])
	}
	// 从chan读出排序好的数据
	for i := 0; i < len(s); i++ {
		s[i] = <-outChan
	}
}

// MergeSort 归并排序 O(nlogn)
// 分治+递归 http://c.biancheng.net/algorithm/merge-sort.html
func MergeSort(s []int) {
	// 合并两个有序子序列 [left, right]
	merge := func(left, mid, right int) {
		// 左序列: [left, mid)  右序列: [mid, right]
		// 需要合并的序列总长度
		length := right - left + 1
		// 用来存合并后的结果
		tmp := make([]int, length)
		// 起点
		leftP := left
		rightP := mid
		tmpP := 0
		for leftP < mid || rightP <= right {
			if leftP < mid && rightP <= right {
				// left和right都还有元素,拿出较小的那个放进tmp
				if s[leftP] < s[rightP] {
					// left较小
					tmp[tmpP] = s[leftP]
					leftP++
				} else {
					// right较小
					tmp[tmpP] = s[rightP]
					rightP++
				}
			} else if leftP < mid {
				// left还有元素
				tmp[tmpP] = s[leftP]
				leftP++
			} else {
				// right还有元素
				tmp[tmpP] = s[rightP]
				rightP++
			}
			tmpP++
		}
		// 把tmp的数据写入到s
		for i := 0; i < len(tmp); i++ {
			s[left+i] = tmp[i]
		}
	}
	// 对区间[left, right]进行合并排序
	var mergeSort func(left, right int)
	mergeSort = func(left, right int) {
		// 相等的情况就是只有一个元素,不做任何操作
		if left != right {
			mid := (right + left) / 2
			// 先把两边排好序
			mergeSort(left, mid)
			mergeSort(mid+1, right)
			// 然后合并两边,两边都是有序子序列
			merge(left, mid+1, right)
		}
	}
	mergeSort(0, len(s)-1)
}
