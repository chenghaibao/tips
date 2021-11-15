package leetcode46

import (
	"fmt"
	"testing"
)

var nums []int
var result [][]int

func init() {
	nums = []int{1, 2, 3}
}
func Test(t *testing.T) {
	fmt.Println(permute(nums))
}

/**
	给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
	这种回溯的要看一下  不难  但是不好写  思路较为难梳理  自我感觉面试出回溯的话  这道题的几率还蛮大的
*/

// 文字讲解
	// 我第一次写了三个for 虽然也可以  但是超时了
	// 然后看了答案   发现回溯，但是不好理解
	// 每一次都想加，然后去判断数字是否已使用过，每个情况都要有


// 挺不好理解的  这种回溯就比较烦人
// ----------
// 第一次循环  【1】 used[true],   [1,2]  used[true,true]  , [1,2,3] [true,true,true]
// 第一次结束  开始第一次里面没有用到的子循环 接着跑
// 子循环推到过程  [1] [true,false,false]  [1,3]  [true false true]  , [1,3,2] [true,true,true]
// 选择，退回原本步骤，中止条件
// -----------

// 回溯核心
// nums: 原始列表
// pathNums: 路径上的数字
// used: 是否访问过
func backtrack(nums, pathNums []int, used []bool) {
	// 结束条件：走完了，也就是路径上的数字总数等于原始列表总数
	if len(nums) == len(pathNums) {
		tmp := make([]int, len(nums))
		// 切片底层公用数据，所以要copy
		copy(tmp, pathNums)
		// 把本次结果追加到最终结果上
		result = append(result, tmp)
		return
	}

	// 开始遍历原始数组的每个数字
	for i := 0; i < len(nums); i++ {
		// 检查是否访问过
		if !used[i] {
			// 没有访问过就选择它，然后标记成已访问过的
			used[i] = true
			// 做选择：将这个数字加入到路径的尾部，这里用数组模拟链表
			pathNums = append(pathNums, nums[i])
			fmt.Println(i,pathNums, used)
			backtrack(nums, pathNums, used)
			// 撤销刚才的选择，也就是恢复操作
			pathNums = pathNums[:len(pathNums)-1]

			// 标记成未使用
			used[i] = false
			fmt.Println(pathNums, used)
		}
	}
}

func permute(nums []int) [][]int {
	var pathNums []int
	var used = make([]bool, len(nums))
	// 清空全局数组（leetcode多次执行全局变量不会消失）
	result = [][]int{}
	backtrack(nums, pathNums, used)
	return result
}
