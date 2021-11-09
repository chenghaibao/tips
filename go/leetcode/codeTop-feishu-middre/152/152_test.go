package leetcode209

import (
	"fmt"
	"hb_leetcode/codeTop-feishu-easy/utils"
	"testing"
)

func TestSpell(t *testing.T) {
	fmt.Println(maxProduct([]int{2, 3, -2,4}))
}

// 暴力
//func maxProduct(nums []int) int {
//	if len(nums) == 0 {
//		return 0
//	}
//	temp := 1
//	res := nums[0]
//	for i := 0; i < len(nums); i++ {
//		temp = 1
//		for j := i; j < len(nums); j++ {
//			temp *= nums[j]
//			if res < temp {
//				res = temp
//			}
//		}
//	}
//	return res
//}


// dp
func maxProduct(nums []int) int {
	maxF, minF, ans := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		maxF = utils.Max(maxF * nums[i], utils.Max(nums[i], minF * nums[i]))
		minF = utils.Min(minF * nums[i],  utils.Min(nums[i], maxF * nums[i]))
		ans = utils.Max(maxF, ans)
	}
	return ans
}