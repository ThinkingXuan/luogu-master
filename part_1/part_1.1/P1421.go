package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	price := 19
	money := a*10 + b
	fmt.Println(money / price)
}
