package main

import "fmt"

//给定一个整数，请将该数各个位上数字反转得到一个新数。新数也应满足整数的常见形式，即除非给定的原数为零，否则反转后得到的新数的最高位数字不应为零（参见样例2）。

func main() {
	var n int
	fmt.Scan(&n)
	flag, res := 1, 0
	if n < 0 {
		flag = -1
		n = -n
	}

	for n > 0 {
		res = res*10 + n%10
		n = n / 10
	}
	fmt.Println(flag * res)
}
