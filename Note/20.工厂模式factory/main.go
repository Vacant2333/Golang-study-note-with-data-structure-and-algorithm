package main

import (
	"Golang-study-note-with-data-structure-and-algorithm/Note/20.工厂模式factory/model"
	"fmt"
)

/*
	golang的结构体没有构造函数,永昌可以使用工厂模式解决这个问题
	当结构体定义的时候首字母是小写的时候但是想在别的包使用的时候则可以使用工厂模式
*/

func main() {
	var stu1 = model.Student1{Name: "123"}
	fmt.Println(stu1)
	// 工厂模式 使用student2(开头小写 私有)
	stu2 := model.NewStudent2("男")
	fmt.Println(stu2)
}
