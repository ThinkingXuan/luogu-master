package main

import (
	"fmt"
	"sort"
)

//https://www.luogu.com.cn/problem/P1059

func main() {
	var n int
	fmt.Scan(&n)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}

	// 排序
	sort.Ints(nums)

	// 去重
	index := 1
	for i := 1; i < n; i++ {
		if nums[i] != nums[i-1] {
			nums[index] = nums[i]
			index++
		}
	}
	nums = nums[:index]

	fmt.Println(index)
	for i := 0; i < index; i++ {
		fmt.Print(nums[i], " ")
	}
}
