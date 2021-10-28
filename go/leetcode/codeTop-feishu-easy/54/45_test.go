package leetcode45

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
		Val: 3,
		Left: &TreeNode{
			Val:  1,
			Left: nil,
			Right: &TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   4,
			Left:  nil,
			Right: nil,
		},
	}
}

// dfs 判断左节点等于右节点   右节点等于子节点
func TestQ(t *testing.T) {
	fmt.Println(kthLargest(root, 1))
}

// 一次插入node 节点的值
func kthLargest(root *TreeNode, k int) int {
	nodeMap := make([]int, 0)
	dfs(root, &nodeMap)
	return nodeMap[k-1]
}

func dfs(root *TreeNode, nodeMap *[]int) {
	if root.Right != nil {
		dfs(root.Right, nodeMap)
	}
	if root != nil  {
		*nodeMap = append(*nodeMap, root.Val)
	}
	if root.Left != nil {
		dfs(root.Left, nodeMap)
	}

}
