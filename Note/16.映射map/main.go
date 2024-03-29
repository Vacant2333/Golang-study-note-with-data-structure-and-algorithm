package main

import (
	"fmt"
	"sort"
)

/*
	map 是 key -> value 数据结构,又称为字段或关联数组 (字典/映射)
	基本语法:
		map[key_type]value_type
		key的类型:
			可以是bool int string pointer channel array struct interface 等等
			通常为int string
			但不能是slice map function,因为他们不能用==判断
		value的类型:
			和key基本一样
			通常: int float string map struct
	map声明之后是不会分配内存的,需要用make初始化 分配内存
	map的key是不能重复的 但是value可以重复
*/

func main() {
	// 声明
	// var m map[string]int
	// key是string 值是 map[string]int
	// var m2 map[string]map[string]int
	m2 := map[string]int{"a": 1}
	fmt.Println(m2)

	m := make(map[string]int, 10)
	m["a1"] = 10
	m["a2"] = 90
	fmt.Println(m)

	// map的删除操作 如果key不存在不会报错
	delete(m, "a1")
	fmt.Println(m)

	// 通过ok-idiom模式判断键值对是否存储
	if value, ok := m["a3"]; ok {
		// 这里不会执行 因为没有a3这个key
		fmt.Println("有这个key value是", value)
	}

	// 实现顺序读取,先sort key然后读
	unsortMap := map[string]int{
		"s1": 2,
		"s3": 25,
		"s2": 1,
		"s0": 50,
	}
	keys := make([]string, 0)
	for k := range unsortMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("k:%v v:%v\n", k, unsortMap[k])
	}
}
