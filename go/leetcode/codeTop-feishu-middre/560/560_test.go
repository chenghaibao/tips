package leetcode560

import (
	"fmt"
	"testing"
)

var nums []int
var k int

func init() {
	nums = []int{1, 1, 1}
	k = 2
}

func Test(t *testing.T) {
	fmt.Println(subarraySum(nums, k))
}

//func subarraySum(nums []int, k int) int {
//	count := 0
//
//	// 遍历数组
//	for start:=0;start<len(nums);start++{
//		sum := 0
//
//		// 遍历子数组
//		for end := start; end >= 0;end--{
//			sum += nums[end]
//
//			// 如果结果相同则总数++
//			if sum == k {
//				count++
//			}
//		}
//	}
//
//	return count
//}

func subarraySum(nums []int, k int) int {
	count := 0
	hash := map[int]int{0: 1}
	preSum := 0

	// 遍历数组
	for i := 0; i < len(nums); i++ {
		// 前缀和
		preSum += nums[i]
		// 前缀和是否存在   前缀和-当前和.是否存在  存在即出现过
		if hash[preSum-k] > 0 {
			count += hash[preSum-k]
		}
		// 添加前缀和
		hash[preSum]++
		fmt.Println(preSum,hash[preSum])
	}
	return count
}
