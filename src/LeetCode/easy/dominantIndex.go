package main

//47. 至少是其他数字两倍的最大数

import (
	"fmt"
	"sort"
)

func dominantIndex(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		return 0
	}
	m := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = i
	}
	sort.Sort(sort.IntSlice(nums))
	if nums[len(nums) - 1] >= nums[len(nums) - 2] * 2 {
		return m[nums[len(nums) - 1]]
	}
	return -1
}

func main() {
	nums := []int{}
	fmt.Println(dominantIndex(nums))
}
