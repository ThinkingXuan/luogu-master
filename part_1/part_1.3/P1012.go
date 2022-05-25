package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	res := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&res[i])
	}

	sort.Slice(res, func(i, j int) bool {
		s1, s2 := fmt.Sprintf("%s%s", res[i], res[j]), fmt.Sprintf("%s%s", res[j], res[i])
		return s1 > s2
	})

	fmt.Println(strings.Join(res, ""))
}
