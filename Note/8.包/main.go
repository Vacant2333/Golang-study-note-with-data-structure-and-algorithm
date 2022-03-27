package main

func main() {

}

/*
	package 一定要在第一行

	import 包的时候 路径从GOPATH/src下开始 不需要带src

	如果要编译成一个可执行程序文件 那这个包必须是main 即package main

	go 的每一个文件都属于一个包 go是以包的形式来管理文件和项目目录结构

	包中 函数/全局变量名 大写字母开头就可以从别的包访问 小写则不行(私有)

	包名和文件夹名 没有任何关系 只是这样比较方便  一个文件夹也可能有多个包
	比如: 文件名叫a 包名叫b,代码中import a 但是使用的时候却是b.xxx

	引入包:
		import "fmt"
		或是
		import (
			"fmt"
			"utils"
		}

	访问其他包的函数/变量时 语法是  包名.函数名

	给包取别名 (原来的名字就不能用了,比如fmt就不能用)
	import (
		f "fmt"
	)
*/
