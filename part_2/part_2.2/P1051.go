package main

import (
	"fmt"
	"sort"
)

// https://www.luogu.com.cn/problem/P1051

func main() {
	var n int
	fmt.Scan(&n)

	res := make([]struct {
		name  string
		money int
	}, n)
	sum := 0
	for i := 0; i < n; i++ {
		tmp := 0 // 每个人的奖学金数额

		var name string
		var score1, score2, cnt int
		var flag1, flag2 string
		fmt.Scan(&name, &score1, &score2, &flag1, &flag2, &cnt)

		if score1 > 80 && cnt >= 1 {
			tmp += 8000
		}
		if score1 > 85 && score2 > 80 {
			tmp += 4000
		}
		if score1 > 90 {
			tmp += 2000
		}
		if score1 > 85 && flag2[0] == 'Y' {
			tmp += 1000
		}
		if score2 > 80 && flag1[0] == 'Y' {
			tmp += 850
		}

		res[i] = struct {
			name  string
			money int
		}{name: name, money: tmp}

		sum += tmp
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i].money > res[j].money
	})

	fmt.Println(res[0].name)
	fmt.Println(res[0].money)
	fmt.Println(sum)
}
