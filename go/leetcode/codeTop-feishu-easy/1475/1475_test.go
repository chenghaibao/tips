package leetcode1475

import (
	"fmt"
	"testing"
)

func TestBaoli(t *testing.T) {
	prices := []int{8, 4, 6, 2, 3}
	fmt.Println(finalPrices(prices))
}

// 暴力解法
func finalPrices(prices []int) []int {
	// 自己写的
	//price := make([]int, 0)
	//count := len(prices)
	//for i := 0; i < count; i++ {
	//	isFound := false
	//	for j := i + 1; j < count; j++ {
	//		if prices[j] < prices[i]{
	//			isFound = true
	//			price = append(price,prices[i]-prices[j])
	//			break
	//		}
	//	}
	//	if !isFound {
	//		price = append(price, prices[i])
	//	}
	//}
	//return price

	// leetcode 的
	for i,p:=range prices{
		for _,q:=range prices[i+1:]{
			if q<=p{
				prices[i]=p-q
				break
			}
		}
	}
	return prices
}

