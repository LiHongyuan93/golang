package main

import "fmt"

/*
1518. 换酒问题
*/

func numWaterBottles(numBottles int, numExchange int) int {
	if numExchange == 0 {
		return numBottles
	}
	drink := numBottles
	empty := numBottles
	for empty/numExchange != 0 {
		drink = drink + empty/numExchange
		empty = empty/numExchange + empty%numExchange
		fmt.Println("drink, empty:", drink, empty)
	}
	return drink
}

func main() {
	numBottles := 4
	numExchange := 0
	fmt.Println(numWaterBottles(numBottles, numExchange))
}
