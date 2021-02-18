package main

import (
	"fmt"
)

/*
258. 各位相加
 */

func addDigits(num int) int {
	var arr []int
	sum := 0
	sign := 0

	fmt.Println("num,",num)
	//sign为标志，sign = 1时退出循环
	for sign == 0 {
		//将整数转化为数组
		arr = append(arr, num % 10)
		num = num / 10

		if num == 0 {
			sum = 0
			for _, v := range arr {
				sum = sum + v
			}
			// 检查各个位上的数字相加结果是否为一位数，是：退出循环，不是：继续循环
			if sum / 10 == 0 {
				sign = 1
			} else {
				num = sum
				arr = []int{}
			}
		}
	}
	return sum
}

func main() {
	num := 101991938392
	fmt.Println(addDigits(num))
}