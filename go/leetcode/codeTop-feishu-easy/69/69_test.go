package leetcode69

import (
	"fmt"
	"testing"
)

func TestMySqrt(t *testing.T){
	fmt.Println(mySqrt(4))
	fmt.Println("big apple")
}


// 确定二分的上下边界   mid*mid <= X  二分查找
func mySqrt(x int)  int {
	l, r := 0, x
	ans := 0
	for l <= r {
		mid := l + (r - l) / 2
		fmt.Println( mid)
		if mid * mid <= x {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ans

}