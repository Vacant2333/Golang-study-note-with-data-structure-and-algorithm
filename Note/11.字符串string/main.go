package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
	Go语言编码方式是UTF-8 一个汉字占3个字节

	string的底层是一个byte数组 因此也可以进行切片处理
	string是不可变的 不能通过str[0] = 'z' 来改变
*/

func main() {
	str := "abcde是"
	// 计算字符串长度
	// 注意: 如果包含汉字的时候需要精确计算字符的个数,必须先将字符串转换为rune类型数组然后使用len
	fmt.Println("str的长度(字节)", len(str))

	// 遍历字符串 同时处理中文
	for index, c := range str {
		fmt.Printf("%d->%c[%T]|", index, c, c)
	}
	fmt.Println()
	// 遍历字符串 第二种 也可以处理中文
	str2 := []rune(str)
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%c|", str2[i])
	}
	fmt.Println()

	// 字符串转整数
	n, err := strconv.Atoi("2002")
	if err != nil {
		fmt.Print("转换错误 \n")
	} else {
		fmt.Printf("%T n=%v \n", n, n)
	}

	// 整数转字符串
	str3 := strconv.Itoa(619)
	fmt.Printf("%T %v \n", str3, str3)

	// 字符串转[]byte
	bt := []byte("niubi")
	fmt.Println(bt)

	// []byte转字符串
	s := string(bt)
	fmt.Println(s)

	// 10进制转别的进制
	j := strconv.FormatInt(123, 8)
	fmt.Printf("123对应的八进制:%v %T \n", j, j)

	// 查找子串
	son := "son"
	father := "who is my son?"
	fmt.Println("包含子串:", strings.Contains(father, son))

	// 统计一个字符串有几个指定的子串
	small := "abc"
	big := "abc def abc ddd abo"
	fmt.Println("子串数量:", strings.Count(big, small))

	// 判断字符串是否相等
	s1 := "abc"
	s2 := "ABC"
	fmt.Println("比较字符串(区分大小写):", s1 == s2)
	fmt.Println("比较字符串(不区分大小写):", strings.EqualFold(s1, s2))

	// 返回子串在字符串中第一次出现的index值(没有的话返回-1)
	fmt.Println(strings.Index("我是真的菜", "的"))
}
