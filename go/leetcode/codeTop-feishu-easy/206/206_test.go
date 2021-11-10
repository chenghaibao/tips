package leetcode206

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

var head *ListNode

func init() {
	head = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
						},
					},
				},
			},
		},
	}
}

func TestDieDai(t *testing.T){
	fmt.Println(reverseList(head))
}

// next 指向下一级，  现在next 指向 pre
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head

	// data next   迭代反转
	for cur != nil {
		// 下一次要循环的链表
		next := cur.Next
		// 链表都是指向下一个指针，所以说我们反转链表只需要指向上一个指针即可
		cur.Next = pre
		pre = cur
		// 重新给予练表
		cur = next
	}

	return pre
}


/**
	文字讲解
		206 -- 反转链表
			正常单链表都是指向下一级的指针，反转链表我们只有做一下替换，也就是指向上一个指针即可
			先设定一个头部指针， 依次指向即可达成题目的需求
 */