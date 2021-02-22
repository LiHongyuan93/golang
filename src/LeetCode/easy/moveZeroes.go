package main

import "fmt"

func moveZeroes(nums []int)  {
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			nums = append(nums[:i],nums[i+1:]...)
			nums = append(nums, 0)
			i--
			n--
		}
	}
	fmt.Println(nums)
}

func main() {
	nums := []int{0,0,1}
	moveZeroes(nums)
}
