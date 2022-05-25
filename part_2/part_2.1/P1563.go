package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	dir := make([]int, n)
	nums := make([]string, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&dir[i], &nums[i])
	}

	res := make([]int, m)
	step := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&res[i], &step[i])
	}

	cur := 0
	for i := 0; i < m; i++ {
		k, v := res[i], step[i]

		if k == dir[cur] {
			cur = (cur + n - v) % n
		} else {
			cur = (cur + v) % n
		}
	}
	fmt.Println(nums[cur])
	return
}
