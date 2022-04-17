package set

/*
	集合:
		没有重复项的一组数据,可以实现Union(合并),Find(查找,返回父节点的下标)

	father(父节点的上标，负数表示它是一个父节点，也就是SetName，他的绝对值就是这个集合的元素个数)
	index  data  father
	0      10     -2
	1      20     1
	2	   30     1
*/

// ElementType 集合的数据类型
type ElementType int

// SetName 集合的名称,根节点的下标作为名称
type SetName int

type Set []Node

type Node struct {
	Data   ElementType
	Father int
}

// SetData 存储所有的集合的节点
var SetData Set

// Create 用一个slice创建一个Set 返回SetName
func Create(s []ElementType) SetName {
	// 父节点
	fNode := insert(s[0], -1)
	for i := 1; i < len(s); i++ {
		insert(s[i], fNode)
	}
	return SetName(fNode)
}

// Union 合并两个集合(将较小的集合的所有元素的Father改为较大的集合的Father),返回新集合的名称
func Union(set1, set2 SetName) SetName {
	if SetData[set1].Father*-1 >= SetData[set2].Father*-1 {
		// 更新节点总数
		SetData[set1].Father += SetData[set2].Father
		// set1的元素比set2多，把set2的元素的Father改为set1
		SetData[set2].Father = int(set1)
		for i, v := range SetData {
			if v.Father == int(set2) {
				SetData[i].Father = int(set1)
			}
		}
		return set1
	} else {
		SetData[set2].Father += SetData[set1].Father
		SetData[set1].Father = int(set2)
		for i, v := range SetData {
			if v.Father == int(set1) {
				SetData[i].Father = int(set2)
			}
		}
		return set2
	}
}

// Find 在集合中查找一个节点，如果没有返回-1，有的话返回父节点的下标
func Find(data ElementType) SetName {
	for i := 0; i < len(SetData); i++ {
		if SetData[i].Data == data {
			return SetName(SetData[i].Father)
		}
	}
	return -1
}

// Insert 插入一个节点,返回下标
func insert(data ElementType, father int) int {
	if father != -1 {
		// 父节点计数+1
		SetData[father].Father--
	}
	SetData = append(SetData, Node{data, father})
	return len(SetData) - 1
}
