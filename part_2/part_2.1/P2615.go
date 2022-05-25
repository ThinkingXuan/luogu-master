package main

import "fmt"

//https://www.luogu.com.cn/problem/P2615

func main() {
	var n int
	fmt.Scan(&n)

	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}

	curi, curj := 0, n/2
	res[curi][curj] = 1

	for k := 2; k <= n*n; k++ {

		if curi == 0 && curj != n-1 {
			curi, curj = n-1, curj+1
		} else if curi != 0 && curj == n-1 {
			curi, curj = curi-1, 0
		} else if curi == 0 && curj == n-1 {
			curi = curi + 1
		} else if curi != 0 && curj != n-1 && res[curi-1][curj+1] == 0 {
			curi, curj = curi-1, curj+1
		} else {
			curi = curi + 1
		}

		res[curi][curj] = k
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(res[i][j], " ")
		}
		fmt.Println()
	}

}
