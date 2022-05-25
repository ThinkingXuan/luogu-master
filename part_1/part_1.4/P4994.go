package main

import "fmt"

//记忆数组
var fp [10000002]int64
var m int64

// 记忆化，斐波那契
func f(i int64) int64 {
	if fp[i] != 0 { //调取记忆
		return fp[i]
	}
	if i == 1 || i == 2 {
		fp[i] = 1 % m
		return fp[i]
	} else {
		fp[i] = (f(i-1) + f(i-2)) % m
		return fp[i]
	}
}

func main() {
	fmt.Scan(&m)
	var i int64 = 1
	for f(i) != 0 || f(i+1) != 1 {
		i++
	}
	fmt.Println(i)
	return
}
