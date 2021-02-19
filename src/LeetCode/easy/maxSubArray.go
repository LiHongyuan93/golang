package main

/*
53. 最大子序和
 */

import "fmt"

func maxSubArray(nums []int) int {
	n := len(nums)
	res := nums[0]
	tmp := 0

	if n == 1 {
		return nums[0]
	}

	// 遍历左下标l
	for l := 0; l < n; l++ {
		tmp = nums[l]
		if tmp > res {
			res = tmp
		}
		// l如果是数组最后一位，循环终止
		if l == n - 1 {
			break
		}
		// 当左下标为l时，遍历右下标
		for r := l + 1; r < n; r++ {
			tmp = tmp + nums[r]
			if tmp > res {
				res = tmp
			}
		}
	}
	return res
}

func main() {
	nums := []int{-1,-2}
	fmt.Println(maxSubArray(nums))
}
