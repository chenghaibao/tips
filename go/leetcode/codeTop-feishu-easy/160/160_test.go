package leetcode160

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

var listA *ListNode
var listB *ListNode

// hashtable 解决方式
func init() {
	listA = &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 8,
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
	listB = &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 8,
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
		},
	}
}

func TestDieDai(t *testing.T) {
	fmt.Println(getIntersectionNode(listA,listB))
}

// next 指向下一级，  现在next 指向 pre
func getIntersectionNode(listA *ListNode, listB *ListNode) *ListNode {
	in := map[*ListNode]bool{}

	for tmp := listA; tmp != nil; tmp = tmp.Next {
		in[tmp] = true
	}


	for tmp := listB; tmp != nil; tmp = tmp.Next {
		if in[tmp]{
			return tmp
		}
	}
	return nil

}
