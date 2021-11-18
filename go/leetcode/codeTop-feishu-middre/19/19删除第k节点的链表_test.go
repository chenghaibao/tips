package _8

import (
	"fmt"
	"testing"
)

/**
	给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
*/


// 文字讲解
// 采取快慢指针操作
// 先让开始节点跑出n节点，那么开始节点结束end节点就等于倒数节点，跳过就好了
// 不好理解的地方就是 为什么先让start节点先跑出n个节点，其实就是start与end之间的距离即n
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

func Test(t *testing.T) {
	fmt.Println(removeNthFromEnd(head, 2))
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 亚节点
	middle := &ListNode{0, head}
	//  start  超前 n 个节点。当 start 遍历到链表的末尾时，end 就恰好处于倒数第n个节点
	start, end := head, middle

	// 确定第second指针从哪里开使
	for i := 0; i < n; i++ {
		start = start.Next
	}

	// 如果开始指针结束  那么当前的指针就是被删除  需要跳过Next(即删除指针)
	for ; start != nil; start = start.Next {
		end = end.Next
		fmt.Println(start.Val,"-----",end.Val)
	}

	// 跳过指针 （内存引用）
	end.Next = end.Next.Next
	return middle.Next
}