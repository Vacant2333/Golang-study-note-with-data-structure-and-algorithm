package model

type Student1 struct {
	Name string
}

// 私有 只能在本包使用
type student2 struct {
	sex string
}

// NewStudent2 通过工程模式从外部使用student2
// 可以理解成这个就是student2类的 构造函数 (共有方法调用私有变量)
func NewStudent2(sexValue string) *student2 {
	return &student2{sex: sexValue}
}
