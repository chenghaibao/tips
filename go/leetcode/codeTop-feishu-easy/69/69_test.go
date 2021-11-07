package leetcode101

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
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left:  &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
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
	fmt.Println(invertTree(root))
}

func invertTree(root *TreeNode) *TreeNode {
	return dfs(root)
}

func dfs(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left   := dfs(root.Left)
	right  := dfs(root.Right)
	root.Right =  left
	root.Left  =  right
	return root
}