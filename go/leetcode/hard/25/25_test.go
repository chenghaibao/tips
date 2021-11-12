package leetcode25

// 抄袭leetcode 答案

// https://leetcode-cn.com/problems/reverse-nodes-in-k-group/solution/k-ge-yi-zu-fan-zhuan-lian-biao-by-leetcode-solutio/
// 请查看视频与代码
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

func Test(t *testing.T) {
	fmt.Println(reverseKGroup(head, 2))
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	// 创造头节点
	init := &ListNode{Next: head}
	pre := init

	// 遍历链表
	for head != nil {
		// 尾节点
		end := pre
		// 遍历
		for i := 0; i < k; i++ {
			// end 就是k的next 节点   也就是是否需要反转的节点值
			end = end.Next

			// 不存在不需要反转
			if end == nil {
				return init.Next
			}
		}
		// next 下一步的链表值
		next := end.Next

		// 反转列表
		head, end = myReverse(head, end)

		// 确定头节点
		pre.Next = head
		// 反转过后确定尾节点
		end.Next = next

		// 返回反转过后  处理后的链表
		pre = end

		// 开始下一次的链表执行
		head = end.Next
	}
	return init.Next
}

// 反转链表 昨天写过
func myReverse(head, end *ListNode) (*ListNode, *ListNode) {
	pre := end.Next
	// 正常情况下  pre == end  也就是不需要反转
	cur := head
	// 链表符合反转 需要反转
	for pre != end {
		// 下一次要循环的链表
		next := cur.Next
		// 链表都是指向下一个指针，所以说我们反转链表只需要指向上一个指针即可
		cur.Next = pre
		pre = cur
		// 重新给予练表
		cur = next
	}
	// 返回反转后的链表
	return end, head
}


/**
文字讲解
	25 -- K 个一组翻转链表
		总结官方思路
		虚拟一个头尾（pre,end）节点, 如果遍历链表的next指针等于k符合反转要求，就行判断next指针不等于next指针/反之不进行处理，
		就开始反转链表并把反转后的end只想nex移动指针，进行链表的闭环，最后返回虚拟节点的next即是分组反转的链表
*/