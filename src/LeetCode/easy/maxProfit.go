package main

import "fmt"

func maxProfit(prices []int) int {
	sum := 0
	input := 0
	output := 0
	flag := 0
	for i := 0; i < len(prices) - 1; i++ {
		// 什么时候可以买
		if prices[i] < prices[i+1] && flag == 0 {
			input = prices[i]
			flag = 1
			fmt.Println(input)
		}
		//什么时候可以卖
		if prices[i] > prices[i+1] && flag == 1 {
			output = prices[i]
			fmt.Println(output)
			sum = sum + (output - input)
			flag = 0
		}
		// 当遍历到最后买入还没卖出时（即一直递增）
		if flag == 1 && i == len(prices) - 2{
			output = prices[len(prices) - 1]
			sum = sum + (output - input)
		}
	}
	return sum
}

func main() {
	prices := []int{1,2,3,4,5}
	fmt.Println(maxProfit(prices))
}
