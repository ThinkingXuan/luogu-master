package main

import (
	"fmt"
	"math"
	"sort"
)

//https://www.luogu.com.cn/problem/P1068

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	type p struct {
		number, score int
	}

	nums := make([]p, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i].number, &nums[i].score)
	}

	// 准备录取人数
	m = int(math.Floor(float64(m) * 1.5))

	sort.Slice(nums, func(i, j int) bool {
		if nums[i].score == nums[j].score {
			return nums[i].number < nums[j].number
		}
		return nums[i].score > nums[j].score
	})

	// 重分处理
	index := m
	for nums[m-1].score == nums[index].score {
		index++
	}
	// 实际录取人数
	m = index

	fmt.Println(nums[m-1].score, m)
	for i := 0; i < m; i++ {
		fmt.Println(nums[i].number, nums[i].score)
	}
}
