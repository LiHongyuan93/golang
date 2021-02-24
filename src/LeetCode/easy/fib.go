package main

import "fmt"

/*
509. 斐波那契数
 */

func fib(n int) int {
	arr := make([]int, n+2)
	arr[0] = 0
	arr[1] = 1
	for i := 2; i <= n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	return arr[n]
}

func main() {
	n := 6
	fmt.Println(fib(n))
}