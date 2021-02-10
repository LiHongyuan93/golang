package main
/*
219. 存在重复元素 II
 */
import "fmt"

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)
	for i, v := range nums {
		if _, ok := m[v]; ok && i - m[v] <= k{
			fmt.Println("i,m[v],k: ", i,m[v],k)
			return true
		} else {
			m[v] = i
		}
	}
	return false
}

func main() {
	nums := []int{0,1,2,0}
	k := 2
	fmt.Println(containsNearbyDuplicate(nums, k))
}
