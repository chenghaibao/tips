package main

import (
	"fmt"
)

// 栈的存储类型
type Item interface {
}

// 栈的存储空间
type ItemStack struct {
	items []Item
}

// 加入队列
func (i *ItemStack) in(t Item) *ItemStack {
	i.items = append(i.items, t)
	return i
}

// 剔除栈
func (i *ItemStack) out() *Item {
	item := i.items[0]
	i.items = i.items[1 : len(i.items)]
	return &item
}


var s ItemStack

// 初始话栈
func initStack() *ItemStack  {
	if s.items == nil{
		s = ItemStack{}
	}
	return &s
}

// 运行
func main() {
	s := initStack()
	s.in(1)
	s.in(2)
	s.in(3)
	fmt.Println(s.items)
	fmt.Println(*s.out())
	fmt.Println(*s.out())
	fmt.Println(*s.out())
	fmt.Println(s.items)
}
