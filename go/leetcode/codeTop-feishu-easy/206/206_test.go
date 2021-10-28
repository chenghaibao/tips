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
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre
}
