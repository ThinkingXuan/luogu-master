package main

import "fmt"

func main() {
	arr := make([]int, 0)
	a := -1
	for true {
		fmt.Scan(&a)
		if a == 0 {
			break
		}
		arr = append(arr, a)
	}

	reverse(arr)
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")
	}
}
func reverse(num []int) {
	n := len(num)
	for i := 0; i < n/2; i++ {
		num[i], num[n-i-1] = num[n-i-1], num[i]
	}
}
