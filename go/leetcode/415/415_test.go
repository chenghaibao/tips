package leetcode415

import (
	"fmt"
	"strconv"
	"testing"
)

func TestBaoli(t *testing.T) {
	num1 := "11"
	num2 := "123"
	fmt.Println(addStrings(num1,num2))
}

func addStrings(num1 string, num2 string) string {
	add := 0
	ans := ""
	len1 := len(num1)-1
	len2 := len(num2)-1

	// add = 数字相加的大约10 ans = 结果
	for i,j := len1,len2; i >= 0 || j >= 0 || add != 0; i,j = i-1,j-1{

		var x,y int

		if i >= 0 {
			x = int(num1[i] - '0')
		}

		if j >= 0 {
			y = int(num2[j] - '0')
		}

		result := x+y+add
		ans = strconv.Itoa(result%10) + ans
		add = result/10
	}
	return ans
}


