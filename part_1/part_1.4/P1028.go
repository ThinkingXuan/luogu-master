package main

import "fmt"

// 1                   			一种  1
// 2  = f(1) + 1      			二 种 2 12
// 3  = f(1) + 1             	二种  3 13
// 4  = f(1) + f(2) + 1    		4种   4 24 14  124
// 5  = f(1) + f(2) + 1         4种   5  15， 25，125
// 6  = f(1) + f(2) + f(3)+ 1   6种
// ...
// n = f(1) + f(2) + f(3) + ... + f(n/2) + 1

func main() {
	var n int
	fmt.Scan(&n)

	var f [1001]int //存每一位数的种类

	for i := 1; i <= n; i++ {
		for j := 1; j <= i/2; j++ {
			f[i] += f[j]
		}
		f[i]++ //加上本身
	}

	fmt.Println(f[n])
}
