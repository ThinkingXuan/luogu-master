package main

import "fmt"

// 一元三次方程求解 https://www.luogu.com.cn/problem/P1024

var a, b, c, d float64

func f(x float64) float64 {
	return a*x*x*x + b*x*x + c*x + d
}

// 题目分析：  x1 <x2时 f(x1) * f(x2) < 0  有根，，，，或者x1, x2本身是根。
// 根范围 -100 100之间， 根之间绝对值大于1,,,

func main() {
	fmt.Scan(&a, &b, &c, &d)
	cnt := 0

	var i float64
	for i = -100; i < 100; i++ {
		left, right := i, i+1
		f1, f2 := f(left), f(right)

		if f1 == 0 {
			fmt.Printf("%.2f ", left)
			cnt++
		}

		if f1*f2 < 0 { // left和right之间有根

			// 二分查找
			for right-left >= 0.001 {
				m := (left + right) / float64(2)
				if f(m)*f(right) <= 0 { // 等于号，代表 m和right有一个为根。
					left = m
				} else {
					right = m
				}
			}

			fmt.Printf("%.2f ", right)
			cnt++
		}

		if cnt == 3 {
			break
		}
	}
}
