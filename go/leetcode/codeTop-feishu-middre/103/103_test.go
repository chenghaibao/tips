package leetcode103

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
			Val:   9,
			Left: &TreeNode{
				Val:   8,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   14,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}
}

// dfs 判断左节点等于右节点   右节点等于子节点
func TestDfs(t *testing.T) {
	fmt.Println(zigzagLevelOrder(root))
}
func zigzagLevelOrder(root *TreeNode) (ans [][]int) {
	if root == nil {
		return
	}
	tree := []*TreeNode{root}

	for level := 0; len(tree) > 0; level++ {
		vals := []int{}
		q := tree
		tree = nil

		for _, node := range q {
			vals = append(vals, node.Val)
			if node.Left != nil {
				// 添加下次循环 参数
				tree = append(tree, node.Left)
			}
			if node.Right != nil {
				// 添加下次循环参数
				tree = append(tree, node.Right)
			}
		}
		// 本质上和层序遍历一样，我们只需要把奇数层的元素翻转即可
		if level%2 == 1 {
			for i, n := 0, len(vals); i < n/2; i++ {
				// len 等于2-1  就等与第二个元素   反转成立
				vals[i], vals[n-1-i] = vals[n-1-i], vals[i]
			}
		}
		ans = append(ans, vals)
	}
	return

}
