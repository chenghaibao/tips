package leetcode518

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	fmt.Println(change(5, []int{1, 2, 5}))
}

func change(amount int, coins []int) int {
	// 定义dp数组
	dp := make([]int, amount+1)
	// 初始化,0大小的背包, 当然是不装任何东西了, 就是1种方法
	dp[0] = 1
	// 遍历物品
	for i := 0; i < len(coins); i++ {
		// 遍历背包
		for j := coins[i]; j <= amount; j++ {
			// 推导公式
			// 每次背包相加
			dp[j] += dp[j-coins[i]]
		}
		fmt.Println(dp)
	}
	return dp[amount]

}
