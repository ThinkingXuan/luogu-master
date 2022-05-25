package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	num := make([][2]int, 3)
	for i := 0; i < len(num); i++ {
		var tmp [2]int
		fmt.Scan(&tmp[0], &tmp[1])
		num[i] = tmp
	}

	money := math.MaxInt32
	for i := 0; i < len(num); i++ {
		cnt := int(math.Ceil(float64(n) / float64(num[i][0])))
		money = min(money, cnt*num[i][1])
	}
	fmt.Println(money)
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
