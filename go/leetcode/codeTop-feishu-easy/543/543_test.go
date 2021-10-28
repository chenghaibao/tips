package leetcode543

import (
	"fmt"
	"hb_leetcode/utils"
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
			Val: 2,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: nil,
			Right: nil,
		},
	}
}

func TestDieDai(t *testing.T){
	max := 0
	fmt.Println(dfs(root,max))
}

// 最大深度   递归下面的left或者right是否存在
func dfs(root *TreeNode,max int) int {
	if root == nil {
		return 0
	}
	left := dfs(root.Left,max)
	right := dfs(root.Right,max)
	max = utils.Max(max,left+right)
	return utils.Max(left,right)+1
}
