package main

import (
	"fmt"
)

func main() {
	for {
		var a, b, c int64
		fmt.Scan(&a, &b, &c)
		if a == -1 && b == -1 && c == -1 {
			break
		}
		fmt.Printf("w(%d, %d, %d) = %d\n", a, b, c, w(a, b, c))
	}
}

var mp map[[3]int64]int64

func w(a, b, c int64) int64 {
	if mp == nil {
		mp = make(map[[3]int64]int64, 0)
	}
	if v, ok := mp[[3]int64{a, b, c}]; ok {
		return v
	}
	if a <= 0 || b <= 0 || c <= 0 {
		mp[[3]int64{a, b, c}] = 1
		return mp[[3]int64{a, b, c}]
	}
	if a > 20 || b > 20 || c > 20 {
		mp[[3]int64{a, b, c}] = w(20, 20, 20)
		return mp[[3]int64{a, b, c}]
	}
	if a < b && b < c {
		mp[[3]int64{a, b, c}] = w(a, b, c-1) + w(a, b-1, c-1) - w(a, b-1, c)
		return mp[[3]int64{a, b, c}]
	}
	mp[[3]int64{a, b, c}] = w(a-1, b, c) + w(a-1, b-1, c) + w(a-1, b, c-1) - w(a-1, b-1, c-1)
	return mp[[3]int64{a, b, c}]
}
