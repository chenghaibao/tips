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
	fmt.Println(getKthFromEnd(head,2))
}

// 双指针
func getKthFromEnd(head *ListNode, k int) *ListNode {
	fast,last := head,head

	for fast != nil && k>0{
		fast = fast.Next
		k--
	}

	for fast != nil {
		fast = fast.Next
		last = last.Next
	}
	return last
}
