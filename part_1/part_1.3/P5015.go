package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	res := 0

	// 带空格string, 也就是按行读取。
	inputRead := bufio.NewReader(os.Stdin)
	line, _, _ := inputRead.ReadLine()

	for i := 0; i < len(line); i++ {
		if line[i] != ' ' {
			res++
		}
	}
	fmt.Println(res)
}
