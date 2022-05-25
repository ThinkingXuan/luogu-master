package main

import "fmt"

//P1035 [NOIP2002 普及组] 级数求和

func main() {
	var k int
	fmt.Scan(&k)
	var sum float64
	n := 1
	for {
		sum += 1 / float64(n)

		if sum > float64(k) {
			fmt.Println(n)
			return
		}
		n++
	}
}
