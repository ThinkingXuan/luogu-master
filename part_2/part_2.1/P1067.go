package main

import "fmt"

func main() {
	var n, tmp int
	fmt.Scan(&n)
	for i := n; i >= 0; i-- {
		fmt.Scan(&tmp)
		if tmp != 0 {
			if i != n && tmp > 0 {
				fmt.Print("+")
			}
			if abs(tmp) > 1 || i == 0 {
				fmt.Print(tmp)
			}
			if tmp == -1 && i != 0 {
				fmt.Print("-")
			}
			if i > 1 {
				fmt.Printf("x^%d", i)
			}
			if i == 1 {
				fmt.Print("x")
			}
		}
	}
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
