package leetcode94

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

func init() {
	root = &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}
}

func TestDfs(t *testing.T) {
	fmt.Println(inorderTraversal(root))
}

func inorderTraversal(root *TreeNode) []int {
	ret := make([]int, 0)
	return dfsTree(root, ret)
}

func dfsTree(node *TreeNode, ret []int) []int {
	if node == nil {
		return ret
	}
	ret = dfsTree(node.Left, ret)
	ret = append(ret, node.Val)
	ret = dfsTree(node.Right, ret)
	return ret
}
