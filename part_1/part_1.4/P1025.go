package main

import "fmt"

func main() {
	res := 0
	var n, k int
	fmt.Scan(&n, &k)

	var dfs func(index int, sum int, count int)
	dfs = func(index int, sum int, count int) {
		if count == k {
			//不能写成 count == k && sum == n
			// 因为 count == k + 1 , sum != n  时 k - count 变成负数，但是无限循环。
			// 或者 	if k <= count {
			//			return
			//		}
			if sum == n {
				res++
			}
			return
		}

		if sum > n || count > k {
			return
		}

		// sum 当前和
		// k - count  还剩几个数
		// sum+i*(k-count)  当前和+剩下个事*当前最小的数 <=n
		for i := index; sum+i*(k-count) <= n; i++ {
			dfs(i, sum+i, count+1)
		}
	}

	dfs(1, 0, 0)
	fmt.Println(res)
}
