package main

/*
9. 回文数
 */
import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	redirect := 0
	origin := x
	for x != 0 {
		redirect = redirect * 10 + x % 10
		x = x /10
	}
	return origin == redirect
}

func main() {
	x := 1
	fmt.Println(isPalindrome(x))
}
