package main

/*
7. 整数反转
 */

import "fmt"

func reverse(x int) int {
	res := 0
	for x != 0 {
		if tmp := int32(res) ; (tmp * 10) / 10 != tmp {
			return -1
		}
		res = res * 10 + x % 10
		x =  x / 10
	}
	return res
}

func main() {
	x := -123
	fmt.Println(reverse(x))
}
