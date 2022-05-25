package main

import "fmt"

//https://www.luogu.com.cn/problem/P1179

func main() {
	low, high := 0, 0
	fmt.Scan(&low, &high)

	res := 0
	for i := low; i <= high; i++ {
		tmp := i
		for tmp > 0 {
			t := tmp % 10
			if t == 2 {
				res++
			}
			tmp = tmp / 10
		}
	}
	fmt.Println(res)
}
