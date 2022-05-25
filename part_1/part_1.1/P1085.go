package main

import "fmt"

func main() {
	res := -1
	val := -1
	for i := 0; i < 7; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		if a+b > 8 && a+b > val {
			res = i
			val = a + b
		}
	}
	fmt.Println(res + 1)
}
