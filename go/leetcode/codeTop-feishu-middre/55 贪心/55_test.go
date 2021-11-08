package leetcode55

import (
	"fmt"
	"testing"
)

func TestJump(t *testing.T) {
	fmt.Println(canJUmp([]int{3, 2, 1, 0, 4}))
}

func canJUmp(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	dp := make([]bool, len(nums))
	dp[0] = true

	// 遍历数组最大位置
	for i := 1; i < len(nums); i++ {
		// 跳跃位置
		for j := i - 1; j >= 0; j-- {
			// 跳跃范围能不能达到
			if dp[j] && nums[j]+j >= i {
				dp[i] = true
				break
			}
		}
	}
	fmt.Println(dp)
	return dp[len(nums)-1]
}
