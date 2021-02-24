package main

/*
628. 三个数的最大乘积
 */

import (
	"fmt"
	"sort"
)

func maximumProduct(nums []int) int {
	n := len(nums)
	if n < 3 {
		return 0
	}
	sort.Ints(nums)
	res := max(nums[0]*nums[1]*nums[2], nums[0]*nums[1]*nums[n-1])
	fmt.Println(res)
	res = max(res, nums[n-1]*nums[n-2]*nums[n-3])
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	nums := []int{1,2,3}
	maximumProduct(nums)
}