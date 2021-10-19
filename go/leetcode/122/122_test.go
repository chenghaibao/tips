package leetcode122

import (
	"fmt"
	"hb_leetcode/utils"
	"testing"
)

var give []int

func init() {
	give = []int{7, 1, 5, 3, 6, 4}
}

// 暴力解
func TestViolence(t *testing.T) {
	maxValue := 0
	dataCount := len(give)

	// 依次比较  取最大值
	for i := 0; i < dataCount; i++ {
		for j := i + 1; j < dataCount; j++ {
			maxValue = utils.Max(maxValue, give[j]-give[i])
		}
	}
	fmt.Println(maxValue)
}

// 贪心解  正常的思路
func Test(t *testing.T) {
	maxValue := 0
	dataCount := len(give)
	for j :=  1; j < dataCount; j++ {
		maxValue += utils.Max(0,give[j]-give[j-1])
	}
	fmt.Println(maxValue)
}

// 动态规划
func TestMaxProfit(t *testing.T) {
	// 获取最小的输出  博取最大的利益
	// dp[i][0] = max(dp[i][0],-give[i])
	// 最大支出  max(dp[i][0] ,-give[i])
	// 最大收益  max(dp[i][1] , dp[i][0] + give[i])
	n := len(give)
	dp := make([][2]int, n)
	dp[0][1] = -give[0]
	for i := 1; i < n; i++ {
		dp[i][0] = utils.Max(dp[i-1][0], dp[i-1][1]+give[i])
		dp[i][1] = utils.Max(dp[i-1][1], dp[i-1][0]-give[i])
	}
	fmt.Println(dp[n-1][0])
}


