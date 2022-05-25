package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	inputRead := bufio.NewReader(os.Stdin)
	s := ""
	for {
		line, _, _ := inputRead.ReadLine()
		s += string(line)
		if strings.Contains(string(line), "E") {
			break
		}
	}
	a, b, count := 0, 0, 0
	for i := 0; i < len(s); i++ {
		count++
		if s[i] == 'W' {
			a++
		} else if s[i] == 'L' {
			b++
		} else { // s[i] == 'E'
			fmt.Printf("%d:%d\n", a, b)
			break
		}

		if (a >= 11 || b >= 11) && Abs(a-b) >= 2 {
			fmt.Printf("%d:%d\n", a, b)
			a = 0
			b = 0
			count = 0
		}
	}
	fmt.Println()

	a, b, count = 0, 0, 0
	for i := 0; i < len(s); i++ {
		count++
		if s[i] == 'W' {
			a++
		} else if s[i] == 'L' {
			b++
		} else { // s[i] == 'E'
			fmt.Printf("%d:%d\n", a, b)
			break
		}

		if (a >= 21 || b >= 21) && Abs(a-b) >= 2 {
			fmt.Printf("%d:%d\n", a, b)
			a = 0
			b = 0
			count = 0
		}
	}
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
