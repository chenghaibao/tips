package builtin

import (
	"fmt"
	"testing"
)

func TestBuiltin(t *testing.T) {
	aa := make([]int, 0, 10)
	aa = append(aa, []int{1, 2, 3, 4}...)
	fmt.Println(aa)
	hashMap := make(map[int]bool, 12)
	hashMap[12] = true
	hashMap[13] = true
	delete(hashMap, 12)
	count := len(aa)
	copy(aa, aa)
	testIntChannel := make(chan int, 12)
	testIntChannel <- 121321
	aaa := <-testIntChannel
	fmt.Println(hashMap, count, aaa)
	close(testIntChannel)
	fmt.Println(aaa)
}
