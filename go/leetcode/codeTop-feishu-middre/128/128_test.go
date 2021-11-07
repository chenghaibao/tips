package leetcode128

import (
	"fmt"
	"testing"
)

var array []int

func init() {
	array = []int{100, 4, 200, 1, 3, 2}
}

func TestLong(t *testing.T) {
	fmt.Println(longestConsecutive(array))
}

func longestConsecutive(nums []int) int {
	// init hashtable
	hashtable := map[int]bool{}
	// set hashtable
	for _, v := range nums {
		hashtable[v] = true
	}
	// maxCount
	maxCount := 0
	// middle variable
	tmpCount := 0
	// range hashtable
	for num := range hashtable {
		// num has child ，continue exec
		if hashtable[num-1] {
			continue
		}
		tmpCount = 0
		// exist hashtable
		for hashtable[num] {
			num++
			tmpCount++
		}

		if tmpCount > maxCount {
			maxCount = tmpCount
		}
	}

	return maxCount
}
