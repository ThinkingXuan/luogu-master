package main

import "fmt"

func main() {
	l, m := 0, 0
	fmt.Scan(&l, &m)
	arr := make([]int, l+1)

	for i := 0; i < m; i++ {
		u, v := 0, 0
		fmt.Scan(&u, &v)

		for j := u; j <= v && j <= l; j++ {
			arr[j] = 1
		}
	}
	res := 0
	for i := 0; i <= l; i++ {
		if arr[i] == 0 {
			res++
		}
	}
	fmt.Println(res)
}
