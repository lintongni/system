package arraylist

import (
	"fmt"
	"testing"
)

func TestArrayList(t *testing.T) {
	l := NewArrayList()
	l.Add(1)
	l.Add(2)
	l.Add(3)
	fmt.Println("******************")
	fmt.Println(l.OrderRead())
	fmt.Println(l.PostRead())
	// 在指定位置插入元素
	fmt.Println("******************")
	l.Insert(9, 2)
	fmt.Println(l.OrderRead())
	fmt.Println(l.PostRead())
	l.Insert(8, 1)
	fmt.Println("******************")
	fmt.Println(l.OrderRead())
	fmt.Println(l.PostRead())
	l.Insert(6, 5)
	fmt.Println("******************")
	fmt.Println(l.OrderRead())
	fmt.Println(l.PostRead())
	l.Insert(11, 7)
	fmt.Println("******************")
	fmt.Println(l.OrderRead())
	fmt.Println(l.PostRead())

	// 获得元素
	fmt.Println(l.Get(0))
	fmt.Println(l.Get(-1))
	fmt.Println(l.Get(-2))
	fmt.Println(l.Get(-8))

	// 删除元素
	l.Remove(-1)
	fmt.Println("******************")
	fmt.Println(l.OrderRead())
	fmt.Println(l.PostRead())
}
