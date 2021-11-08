package leetcode139

import (
	"fmt"
	"testing"
)

func TestWordBeak(t *testing.T) {
	s := "leetcode"
	wordDict := []string{"leet", "code"}
	fmt.Println(wordBreak(s, wordDict))
}

func wordBreak(s string, wordDict []string) bool {
	wordDictSet := make(map[string]bool)

	// hasnMap init
	for _, w := range wordDict {
		wordDictSet[w] = true
	}
	// 确定dp空间
	dp := make([]bool, len(s)+1)
	dp[0] = true

	// 循环字符
	for i := 1; i <= len(s); i++ {
		// 剪枝字符
		for j := 0; j < i; j++ {
			// hash是否存在
			if dp[j] && wordDictSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	// 最后是否剪枝完毕
	return dp[len(s)]
}