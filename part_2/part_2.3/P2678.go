package main

import (
	"fmt"
)

//https://www.luogu.com.cn/problem/P2678
// 从答案入手， 对答案的范围进行二分搜索，然后判断但是是否满足
var nums [50010]int
var l, n, m, ans int

func judge(x int) bool {
	tot := 0 //搬走的石块数量
	i := 0   // 找的石头
	now := 0 // 我站在哪一块石头上
	for i < n+1 {
		i++
		if nums[i]-nums[now] < x { //如果距离小于我二分的答案x，那么这块石头需要搬走
			tot++
		} else {
			now = i // //不然我就跳过来
		}
	}

	if tot > m { //如果要搬走的石头多于m块，那么这个解太大了
		return false
	}
	return true
}

func main() {
	fmt.Scan(&l, &n, &m)
	for i := 1; i <= n; i++ {
		fmt.Scan(&nums[i])
	}

	nums[n+1] = l

	left, right := 1, l

	for left <= right {
		mid := (left + right) / 2
		if judge(mid) {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	fmt.Println(ans)
}
