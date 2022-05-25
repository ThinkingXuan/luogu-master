package main

import "fmt"

//P1980 [NOIP2013 普及组] 计数问题
//试计算在区间 1 到 n 的所有整数中，数字 x（0<x≤9）共出现了多少次？例如，在 11 到 1111 中，即在 1,2,3,4,5,6,7,8,9,10,111,2,3,4,5,6,7,8,9,10,11 中，数字 1 出现了 4 次。
//  暴力方法
func main() {
	var n, x int
	fmt.Scan(&n, &x)
	res := 0
	for i := 1; i <= n; i++ {
		res += helper(i, x)
	}
	fmt.Println(res)
}

func helper(n int, x int) int {
	res := 0
	for n > 0 {
		if n%10 == x {
			res++
		}
		n = n / 10
	}
	return res
}
