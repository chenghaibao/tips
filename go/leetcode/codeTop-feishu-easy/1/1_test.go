package leetcode1

import (
	"fmt"
	"testing"
)

func TestBaoli(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(baoli(nums,target))
}

func baoli(nums []int, target int) []int {
	for i, v := range nums {
		for k, value := range nums {
			if v+value == target {
				return []int{i, k}
			}
		}
	}
	return nil
}


func TestHantable(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i,v:= range nums{
		if p,ok := hashTable[target-v];ok{
			return []int{p,i}
		}
		hashTable[v] = i
	}
	return nil
}

func TestZhizhen(t *testing.T){
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(zhizhen(nums,target))
}

func zhizhen(nums []int, target int) []int {
	count := len(nums)
	i :=0
	for {
		if nums[i] + nums[count-1] == target{
			return []int{i,count-1}
		} else if nums[i] + nums[count-1] > target {
			count--
		} else{
			i++
		}
		if count == 0 && i == count{
			return []int{}
		}
	}
}


