package main

import (
	"encoding/json"
	"fmt"
)

/*
JSON传入Go解析后,int/float都会被认为是float64,如果要转int需要先转为float64后再转为int
*/

type DemoStruct struct {
	Name        string
	Description string
	Num         int
}

func main() {
	jsonStr := "{\"name\": \"JsonT\",\"description\": \"介绍\",\"num\": 666\n}"

	// 第一种 通过一种固定的结构来解析JSON
	var result1 DemoStruct
	json.Unmarshal([]byte(jsonStr), &result1)

	// 第二种 interface大法
	var result2 map[string]interface{}
	json.Unmarshal([]byte(jsonStr), &result2)

	fmt.Printf("result1:%v result2:%v\n", result1, result2)

	// 使用第二种时num会被解析为float64
	fmt.Printf("result2[num]: %v %T\n", result2["num"], result2["num"])
}
