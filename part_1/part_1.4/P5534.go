package main

import "fmt"

func main() {
	var a1, a2, n int

	fmt.Scan(&a1, &a2, &n)
	d := a2 - a1
	an := a1 + (n-1)*d
	sum := ((a1 + an) * n) / 2
	fmt.Println(sum)
}
