package main

/*
面试题 01.01. 判定字符是否唯一
 */

import "fmt"

func isUnique(astr string) bool {
	m := make(map[int32] bool)
	for _, v := range astr {
		if _, has := m[v]; has{
			return false
		}
		m[v] = true
	}
	return true
}

func main() {
	astr := ""
	fmt.Println(isUnique(astr))
}
