package main

import "fmt"

//771. 宝石与石头

func numJewelsInStones(jewels string, stones string) int {
	jewelsCount := 0
	m := make(map[byte]struct{})

	for i := range jewels {
		m[jewels[i]] = struct {}{}
	}

	for j := range stones {
		if _, has := m[stones[j]]; has {
			jewelsCount++
		}
	}
	return jewelsCount
}

func main() {
	J := "z"
	S := "ZZ"
	fmt.Println(numJewelsInStones(J, S))
}