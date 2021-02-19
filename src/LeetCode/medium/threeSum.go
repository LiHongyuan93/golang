package main

/*
15. 三数之和
 */

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	res := [][]int{}
	n := len(nums)
	if len(nums) < 3 {
		return res
	}
	sort.Sort(sort.IntSlice(nums))
	//两层循环
	//先找a
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l := i + 1
		r := n - 1
		if nums[i] > 0 || l >= r {
			return res
		}

		for l < r {
			if nums[i] + nums[l] + nums[r] == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				//判断l r是否有重复
				for nums[l] == nums[l+1] {
					l++
					//防止溢出
					if l == n - 1 {
						break
					}
				}
				for r >= 0 && nums[r] == nums[r-1] {
					r--
					//防止溢出
					if r == 0 {
						break
					}
				}
				l++
				r--
			} else if nums[i] + nums[l] + nums[r] > 0 {
				r--
			} else {
				l++
			}

		}
	}

	return res
}

func main() {
	nums := []int{0,0,0,0,1,-1}
	fmt.Println(threeSum(nums))
}
