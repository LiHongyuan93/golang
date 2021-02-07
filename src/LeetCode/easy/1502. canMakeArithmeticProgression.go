package main

/*
1502. 判断能否形成等差数列
*/
import (
	"fmt"
	"sort"
)

func canMakeArithmeticProgression(arr []int) bool {
	if len(arr) < 3 {
		return true
	}
	sort.Ints(arr)
	for i := 1; i < len(arr)-1; i++ {
		if arr[i]*2 != arr[i-1]+arr[i+1] {
			return false
		}
	}
	return true
}

func main() {
	arr := []int{1, 2, 4}
	fmt.Println(canMakeArithmeticProgression(arr))
}
