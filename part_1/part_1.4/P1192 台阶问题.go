package main

import "fmt"

// f(i) = f(i) + f(i-j)
func main() {
	var n, k int
	fmt.Scan(&n, &k)
	dp := make([]int64, n+1)
	dp[0], dp[1] = 1, 1

	for i := 2; i <= n; i++ {
		for j := 1; j <= k; j++ {
			if i >= j {
				dp[i] = (dp[i] + dp[i-j]) % 100003
			}
		}
	}

	fmt.Println(dp[len(dp)-1])

}
