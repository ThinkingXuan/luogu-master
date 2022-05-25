package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

//hello world.
//aaabbbb
//x
//EOF
//heelo world.
//aaacbbbb
//y<x
//EOF
//60

func main() {
	res := 0
	flag := 0
	var source []string
	var my []string

	time := 0

	inputRead := bufio.NewReader(os.Stdin)
	for {
		line, _, _ := inputRead.ReadLine()
		if string(line) == "EOF" {
			flag++
			if flag != 2 {
				continue
			}
		}

		if flag == 0 {
			source = append(source, string(line))
		}

		if flag == 1 {
			// 去除开头的<
			for line[0] == '<' {
				line = line[1:]
			}
			// 去重中间的<
			newLine := ""
			for i := 0; i < len(line); i++ {
				if line[i] == '<' && len(newLine) > 0 {
					newLine = newLine[:len(newLine)-1]
				} else if line[i] != '<' {
					newLine += string(line[i])
				}
			}
			my = append(my, newLine)
		}
		if flag == 2 {
			line, _, _ = inputRead.ReadLine()
			time, _ = strconv.Atoi(string(line))
			break
		}
	}

	for i := 0; i < len(source) && i < len(my); i++ {
		for j := 0; j < len(source[i]) && j < len(my[i]); j++ {
			if source[i][j] == my[i][j] {
				res++
			}
		}
	}
	res = res * 60
	fmt.Println(Round(float64(res)/float64(time), 0))
}

func Round(val float64, precision int) float64 {
	p := math.Pow10(precision)
	return math.Floor(val*p+0.5) / p
}
