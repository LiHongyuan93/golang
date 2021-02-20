package main

/*
349. 两个数组的交集
 */

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	m1 := make(map[int]int)
	m2 := make(map[int]int)
	res := []int{}
	for _, v := range nums1 {
		if _, has := m1[v]; has {
			continue
		} else {
			m1[v]++
		}
	}

	for _, v := range nums2 {
		if _, has := m2[v]; has {
			continue
		} else {
			m2[v]++
		}
	}
	for key, _ := range m1 {
		if _, has := m2[key]; has {
			res = append(res, key)
		}
	}
	return res
}

func main() {
	nums1 := []int{}
	nums2 := []int{9,4,9,8,4}
	fmt.Println(intersection(nums1, nums2))
}
