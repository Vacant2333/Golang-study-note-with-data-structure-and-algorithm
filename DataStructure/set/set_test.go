package set

import (
	"fmt"
	"testing"
)

func TestCreate(t *testing.T) {
	set1 := Create([]ElementType{1, 2, 3, 4, 5})
	fmt.Println(set1)
}

func TestUnion(t *testing.T) {
	set1 := Create([]ElementType{1, 2, 3, 4, 5})
	set2 := Create([]ElementType{6, 7, 8})
	set3 := Union(set1, set2)
	fmt.Println(set3)
}

func TestFind(t *testing.T) {
	set1 := Create([]ElementType{1, 2, 3, 4, 5})
	fmt.Println(set1, Find(3))
}
