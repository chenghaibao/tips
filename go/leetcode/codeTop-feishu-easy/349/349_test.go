package leetcode349

import (
	"fmt"
	"sort"
	"testing"
)

var num1 []int
var num2 []int

func init() {
	num1 = []int{1, 2, 2, 1}
	num2 = []int{2, 2}
}
func TestBao(t *testing.T) {
	fmt.Println(intersection(num1, num2))
}

func intersection(nums1 []int, nums2 []int) (res []int) {
	// leetcode
	//m := make(map[int]int)
	//for _, v := range nums1 {
	//	m[v] = 1
	//}
	//// 利用count>0，实现重复值只拿一次放入返回结果中
	//for _, v := range nums2 {
	//	if count, ok := m[v]; ok && count > 0 {
	//		res = append(res, v)
	//		m[v]--
	//	}
	//}
	//return res

	// leetcode 双指针
	sort.Ints(nums1)
	sort.Ints(nums2)
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		x, y := nums1[i], nums2[j]
		if x == y {
			if res == nil || x > res[len(res)-1]{
				res = append(res, x)
			}
			i++
			j++
		} else if x < y {
			i++
		} else {
			j++
		}
	}

	return

	// 自己的方案
	//inter := make(map[int]int, 0)
	//for _, j := range nums1 {
	//	for _, v := range nums2 {
	//		if j == v {
	//			inter[j] = v
	//		}
	//	}
	//}
	//for _, j := range inter {
	//	res = append(res,j)
	//}
	//fmt.Println(inter)
	//return
}
