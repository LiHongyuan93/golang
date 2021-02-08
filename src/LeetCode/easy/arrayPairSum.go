package main

import (
	"fmt"
	"sort"
)

/*
561. 数组拆分 I
 */

func arrayPairSum(nums []int) int {
	sum := 0
	sort.Ints(nums)
	for i := 0; i < len(nums); i += 2 {
		sum = sum + nums[i]
	}
	return sum
}

func main() {
	nums := []int{1,2,2,5,6,6}
	fmt.Println(arrayPairSum(nums))
}