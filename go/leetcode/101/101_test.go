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
	fmt.Println(isSymmetric(root))
}

func isSymmetric(root *TreeNode) bool {
	return isTree(root.Left,root.Right)
}

func isTree (left *TreeNode, right *TreeNode)bool{
	if (left == nil || right == nil){
		if(left != nil || right != nil){
			return false
		}
		return true
	}
	if(left.Val == right.Val){
		if(isTree(left.Left,right.Right)&&isTree(left.Right,right.Left)){
			return true
		}
	}
	return false
}