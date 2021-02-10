package main
/*
217. 存在重复元素
 */
import "fmt"

func containsDuplicate(nums []int) bool {
	m := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			return true
		} else {
			m[nums[i]] = true
		}
	}
	return false
}

func main() {
	nums := []int{1,1,1,3,3,4,3,2,4,2}
	fmt.Println(containsDuplicate(nums))
}
