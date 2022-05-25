package main

import "fmt"

func main() {
	n := 0
	fmt.Scan(&n)
	arr := make([]int, n)
	mp := map[int]int{}
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
		mp[arr[i]]++
	}
	res := 0
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			a, b := arr[i], arr[j]
			if v := mp[a+b]; v > 0 {
				res += v
				mp[a+b] = 0
			}
		}
	}
	fmt.Println(res)

}
