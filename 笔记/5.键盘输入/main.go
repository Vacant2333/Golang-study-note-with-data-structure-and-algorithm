package main

import "fmt"

func main() {
	var i int
	var f float64

	// 类似c的scanf
	fmt.Scanf("%d %f", &i, &f)
	fmt.Println(i, f)

	// 按空格分开读取
	fmt.Scan(&i, &f)
	fmt.Println(i, f)
}
