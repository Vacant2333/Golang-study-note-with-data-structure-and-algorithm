package main

/*
请你设计并实现一个满足LRU (最近最少使用) 缓存 约束的数据结构。
实现 LRUCache 类：
LRUCache(int capacity) 以 正整数 作为容量capacity 初始化 LRU 缓存
int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
void put(int key, int value)如果关键字key 已经存在，则变更其数据值value ；
如果不存在，则向缓存中插入该组key-value 。如果插入操作导致关键字数量超过capacity
，则应该 逐出 最久未使用的关键字。
函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。
*/

type node struct {
	last *node
	next *node
	key  int
	val  int
}

type list struct {
	head *node
	tail *node
}

type LRUCache struct {
	capacity int
	l        *list
	m        map[int]*node
}

func Constructor(capacity int) LRUCache {
	l := &list{
		head: &node{},
		tail: &node{},
	}
	l.head.next = l.tail
	l.tail.last = l.head
	return LRUCache{
		capacity: capacity,
		l:        l,
		m:        map[int]*node{},
	}
}

func (this *LRUCache) Get(key int) int {
	if this.exists(key) {
		this.moveToHead(key)
		return this.m[key].val
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if this.exists(key) {
		this.m[key].val = value
		this.moveToHead(key)
	} else {
		if len(this.m) == this.capacity {
			this.deleteLast()
		}
		this.addToHead(key, value)
	}
}

// 添加某个key到头部
func (this *LRUCache) addToHead(key int, value int) {
	this.m[key] = &node{
		key: key,
		val: value,
	}
	this.addNodeToHead(this.m[key])
}

func (this *LRUCache) addNodeToHead(newNode *node) {
	firstTmp := this.l.head.next
	this.l.head.next = newNode
	newNode.last = this.l.head
	newNode.next = firstTmp
	firstTmp.last = newNode
}

// 移动某个key到头部
func (this *LRUCache) moveToHead(key int) {
	this.deleteNode(this.m[key])
	this.addNodeToHead(this.m[key])
}

// 删除最久没用过的
func (this *LRUCache) deleteLast() {
	delete(this.m, this.l.tail.last.key)
	this.deleteNode(this.l.tail.last)
}

// 删除节点
func (this *LRUCache) deleteNode(node *node) {
	node.last.next = node.next
	node.next.last = node.last
}

func (this *LRUCache) exists(key int) bool {
	_, exists := this.m[key]
	return exists
}
