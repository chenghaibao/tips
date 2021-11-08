package leetcode233

import (
	"fmt"
	"math"
	"testing"
)

var Coins []int
var amout int

func init() {
	Coins = []int{1, 2, 5}
	amout = 11
}

func TestTest(t *testing.T) {
	fmt.Println(coinChange(Coins, amout))
}
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	// 初始化dp[0]
	dp[0] = 0
	// 遍历背包,从1开始
	for j := 1; j <= amount; j++ {
		// 初始化为math.MaxInt32
		dp[j] = math.MaxInt32
		// 遍历物品
		for i := 0; i < len(coins); i++ {
			if j >= coins[i] && dp[j-coins[i]] != math.MaxInt32 {
				// 推导公式
				dp[j] = min(dp[j], dp[j-coins[i]]+1)
				//fmt.Println(dp)
			}
		}
	}
	// 没找到能装满背包的, 就返回-1
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

func min(a int, b int) int {
	if a > b {
		return b
	}else {
		return a
	}
}

