package leetcode121

import (
	"fmt"
	"hb_leetcode/utils"
	"math"
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
	// 获取最大int64的值
	minValue := math.MaxInt64
	maxValue := 0

	for _, v := range give {
		// 判断当前循环值是不是比最小的值还要小
		if v < minValue {
			minValue = v
		} else if v-minValue > maxValue {
			// 获取利润
			maxValue = v - minValue
		}
	}
	fmt.Println(maxValue)
}

// 动态规划
func TestMaxProfit(t *testing.T) {
	// 获取最小的输出  博取最大的利益
	// dp[i][0] = max(dp[i][0],-give[i])
	// 最大支出  max(dp[i][0] ,-give[i])
	// 最大收益  max(dp[i][1] , dp[i][0] + give[i])

	dataCount := len(give)
	if dataCount == 0 {
		fmt.Println(0)
	}
	dp := make([][]int, dataCount)
	for i := 0; i < dataCount; i++ {
		dp[i] = make([]int, 2)
	}

	// 第一天买入的价格
	dp[0][0] = -give[0]
	// 第一天的手里的钱
	dp[0][1] = 0
	for i := 1; i < dataCount; i++ {
		//支出
		dp[i][0] = utils.Max(dp[i-1][0], -give[i])
		//收益
		dp[i][1] = utils.Max(dp[i-1][1], dp[i-1][0]+give[i])
	}
	fmt.Println(dp[dataCount-1][1])
}
