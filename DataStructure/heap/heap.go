package heap

/*
	Heap堆(最大堆):
		堆序性:大根堆的左右节点都比根小,也就是头节点是最大值(还有小根堆)
		节点下标为i时，左子节点下标为2i+1,右子节点下标为2i+2,父节点是i/2(3/2 = 2/2 = 1,整除)
		所有的叶子应处于第h/h-1层(h是树的高度),也就是说应是一棵完全二插树
*/
type ElementType int
type Heap []ElementType

// Create 创建一个Heap堆
func Create() Heap {
	return make(Heap, 0)
}

// Insert 插入一个节点在Heap的尾部,然后Shift Up(上滤)最后一个元素,保证堆序性
func (heap *Heap) Insert(data ElementType) {
	*heap = append(*heap, data)

}

// shiftUp 上滤,将一个元素移动到正确的位置(最大堆)
func (heap *Heap) shiftUp(index int) {

}

// swap 交换两个元素
func (heap *Heap) swap(index1, index2 int) {

}
