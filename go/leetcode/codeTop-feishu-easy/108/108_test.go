package leetcode108

import (
	"fmt"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TestMiddle(t *testing.T) {
	fmt.Println(sortedArrayToBST([]int{-10,-3,0,5,9}))
}

func sortedArrayToBST(nums []int) *TreeNode {
	return helper(nums, 0, len(nums)-1)
}

func helper(nums []int, left, right int) *TreeNode {

	// 不会写  参考leetcode  方案 因为数组是排序的  从中间开始取   左边的是左子树，右边的是右子树  中序遍历即可
	if left > right {
		return nil
	}

	// 总是选择中间位置右边的数字作为根节点
	mid := (left + right + 1) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = helper(nums, left, mid-1)
	root.Right = helper(nums, mid+1, right)
	return root
}
