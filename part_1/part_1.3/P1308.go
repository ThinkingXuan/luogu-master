package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	inputRead := bufio.NewReader(os.Stdin)
	line, _, _ := inputRead.ReadLine()

	s := strings.ToLower(string(line))

	line2, _, _ := inputRead.ReadLine()
	t := strings.ToLower(string(line2))

	arrt := strings.Split(t, " ")

	oneA, res := 0, 0

	for i := 0; i < len(arrt); i++ {
		if arrt[i] == s {
			res++
		}
	}

	if res == 0 {
		fmt.Println("-1")
		return
	}
	if s == t[:len(s)] {
		oneA = 0
	} else {
		oneA = strings.Index(t, " "+s+" ") + 1
	}
	fmt.Println(res, oneA)
}
