package leetcode105

import (
	"fmt"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var pre []int
var midder []int

func init() {
	pre = []int{3, 9, 20, 15, 7}
	midder = []int{9, 3, 15, 20, 7}
}

func Test(t *testing.T) {
	fmt.Println(buildTree(pre, midder))
}

func buildTree(preorder []int, inorder []int) *TreeNode {

	if len(preorder) == 0 {
		return nil
	}

	// 每次的节点值
	root := &TreeNode{preorder[0], nil, nil}


	// 每次中续遍历的值
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	// 遍历做节点
	root.Left = buildTree(preorder[1:len(inorder[:0])+1], inorder[:0])
	root.Right = buildTree(preorder[len(inorder[:0])+1:], inorder[0+1:])
	return root
}
