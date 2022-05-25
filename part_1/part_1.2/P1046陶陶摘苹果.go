package main

import "fmt"

func main() {
	var num [10]int
	var high int
	var res int
	for i := 0; i < 10; i++ {
		fmt.Scan(&num[i])
	}
	fmt.Scan(&high)
	all := high + 30
	for i := 0; i < 10; i++ {
		if num[i] <= all {
			res++
		}
	}
	fmt.Println(res)

}
