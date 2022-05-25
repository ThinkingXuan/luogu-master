package main

import (
	"fmt"
	"math"
)

func main() {
	var n, k, res int
	fmt.Scan(&n, &k)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	var dfs func(index int, count, sum int)
	dfs = func(index int, count int, sum int) {
		if count == k && isPrime(sum) {
			res++
		}
		for i := index; i < len(arr); i++ {
			dfs(i+1, count+1, sum+arr[i])
		}
	}
	dfs(0, 0, 0)
	fmt.Println(res)
}

// 埃氏筛写法 (效率更高)  值为-1为质数，其他不是质数
//func sieveMu() []int {
//	const mx int = 1000001
//	//0 1 -1 -1 0 -1 1 -1 0 0 1 -1 0 -1 1 1
//	mu := [mx + 1]int{1: 1} // int8
//	for i := 1; i <= mx; i++ {
//		for j := i * 2; j <= mx; j += i {
//			mu[j] -= mu[i]
//		}
//	}
//	return mu[:]
//}

func isPrime(num int) bool {
	i := 0
	for i = 2; i <= int(math.Floor(math.Sqrt(float64(num)))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
