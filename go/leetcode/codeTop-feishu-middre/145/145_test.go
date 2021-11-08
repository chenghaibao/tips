package leetcode145

import (
	"fmt"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var root *TreeNode

var res []int

func init() {
	root = &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val:  2,
			Left: nil,
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
	}
}

// dfs 判断左节点等于右节点   右节点等于子节点
func TestDfs(t *testing.T) {
	fmt.Println(postorderTraversal(root))
}

func postorderTraversal(root *TreeNode) []int {
	return last(root)
}

// 左右根
func last(node *TreeNode) []int {
	if node == nil {
		return nil
	}
	last(node.Left)
	last(node.Right)
	res = append(res, node.Val)
	return res
}
