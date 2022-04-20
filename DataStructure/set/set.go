package set

/*
	集合:
		没有重复项的一组数据,可以实现Union(合并),Find(查找,返回父节点的下标)

	father(父节点的上标，负数表示它是一个父节点，也就是Name，他的绝对值就是这个集合的元素个数)
	index  data  father
	0      10     -2
	1      20     1
	2	   30     1
*/

// ElementType 集合的数据类型
type ElementType int

// Data 存储所有的集合的节点
var Data []Node

// Name 集合的名称,根节点的下标作为名称
type Name int

// Node 节点
type Node struct {
	Data   ElementType
	Father Name
}

// Create 用一个slice创建一个Set 返回Name
func Create(s []ElementType) Name {
	// 父节点
	fNode := insert(s[0], -1)
	for i := 1; i < len(s); i++ {
		insert(s[i], Name(fNode))
	}
	return Name(fNode)
}

// Union 合并两个集合(将较小的集合的所有元素的Father改为较大的集合的Father),返回新集合的名称
func Union(set1, set2 Name) Name {
	if Data[set1].Father*-1 >= Data[set2].Father*-1 {
		// 更新节点总数
		Data[set1].Father += Data[set2].Father
		// set1的元素比set2多，把set2的元素的Father改为set1
		Data[set2].Father = set1
		for i, v := range Data {
			if v.Father == set2 {
				Data[i].Father = set1
			}
		}
		return set1
	} else {
		Data[set2].Father += Data[set1].Father
		Data[set1].Father = set2
		for i, v := range Data {
			if v.Father == set1 {
				Data[i].Father = set2
			}
		}
		return set2
	}
}

// Find 在集合中查找一个节点，如果没有返回-1，有的话返回父节点的下标
func Find(data ElementType) Name {
	for i := 0; i < len(Data); i++ {
		if Data[i].Data == data {
			if Data[i].Father >= 0 {
				// 它是一个子节点 返回父节点
				return Data[i].Father
			} else {
				// 它是一个父节点 直接返回
				return Name(i)
			}
		}
	}
	return -1
}

// Insert 插入一个节点,返回下标
func insert(data ElementType, father Name) int {
	if father != -1 {
		// 父节点计数+1
		Data[father].Father--
	}
	Data = append(Data, Node{data, father})
	return len(Data) - 1
}
