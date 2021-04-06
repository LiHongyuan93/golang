package main

import (
	"fmt"
	"strconv"
)

func calPoints(ops []string) int {
	stack := make([]int, 0, len(ops))

	for i := range ops {
		switch ops[i] {
		default:
			num, _ := strconv.Atoi(ops[i])
			stack = append(stack, num)
		case "C":
			stack = stack[:len(stack) - 1]
		case "D":
			// <<1 等同于 stack[len(stack)-1]*2
			stack = append(stack, stack[len(stack)-1]<<1)
		case "+":
			stack = append(stack, stack[len(stack) - 1] + stack[len(stack) - 2])
		}
	}

	res := 0
	for  _, v := range stack {
		res = res + v
	}

	return res
}

func main() {
	ops := []string{"5","-2","4","C","D","9","+","+"}
	fmt.Println(calPoints(ops))
}
