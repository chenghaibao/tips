package leetcode45

import (
	"fmt"
	"hb_leetcode/codeTop-feishu-easy/utils"
	"testing"
)


func TestSpell(t *testing.T) {
	fmt.Println(jump([]int{2,3,1,1,4}))
}

func jump(nums []int) int {
	length := len(nums)
	end := 0  // 上次跳跃可达范围右边界（下次的最右起跳点）
	maxPosition := 0 // 目前能跳到的最远位置
	steps := 0
	for i := 0; i < length - 1; i++ {
		fmt.Println(maxPosition,i + nums[i],end)
		maxPosition = utils.Max(maxPosition, i + nums[i])

		// 到达上次跳跃能到达的右边界了
		if i == end {
			fmt.Println("++++++++")
			end = maxPosition // 目前能跳到的最远位置变成了下次起跳位置的有边界
			steps++  // 进入下一次跳跃
		}
	}
	return steps
}



