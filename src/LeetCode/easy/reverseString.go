package main

/*
344. 反转字符串
 */
import "fmt"

func reverseString(s []byte)  {
	for left, right := 0, len(s) - 1; left < right; left++ {
		s[left], s[right] = s[right], s[left]
		right--
	}
}

func main() {
	s := []byte{'h','e','l','l','o'}
	fmt.Println(s)
	reverseString(s)
}
