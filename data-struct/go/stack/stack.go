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

// 加入栈
func (i *ItemStack) push(t Item) *ItemStack {
	i.items = append(i.items, t)
	return i
}

// 剔除栈
func (i *ItemStack) pop() *Item {
	item := i.items[len(i.items)-1]
	i.items = i.items[0 : len(i.items)-1]
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
	s.push(1)
	s.push(2)
	s.push(3)
	fmt.Println(s.items)
	fmt.Println(*s.pop())
	fmt.Println(*s.pop())
	fmt.Println(*s.pop())
	fmt.Println(s.items)
}
