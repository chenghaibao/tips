package leetcode3

import (
	"fmt"
	"strings"
	"testing"
)

func TestStrLen(t *testing.T) {
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}

func lengthOfLongestSubstring(s string) int {
	// 自己写的  有问题   案列倒是能过
	//if s == "" {
	//	return 0
	//}
	//strMaxLen := ""
	//middreStr := ""
	//stop := false
	//hashTable := map[string]bool{}
	//for _, v := range s {
	//	if hashTable[string(v)]  {
	//		if len(middreStr) > len(strMaxLen) {
	//			strMaxLen =  middreStr
	//		}
	//		stop = true
	//		middreStr = string(v)
	//	} else {
	//		if stop == false {
	//			strMaxLen = strMaxLen + string(v)
	//		}
	//		middreStr = middreStr + string(v)
	//	}
	//	hashTable[string(v)] = true
	//}
	//
	//return len(strMaxLen)

	// 滑动窗口
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		// 获取当前子串在的位置
		index := strings.Index(s[start:i], string(s[i]))
		fmt.Println(index)

		// 不存在 结束指针向后移
		if index == -1 {
			if i+1 > end {
				end = i + 1
			}
		} else {
			// 左窗口 向前移 end向前移
			start += index + 1
			end += index + 1
		}
	}

	// 尾 - 开始  等最大字串
	return end - start
}
