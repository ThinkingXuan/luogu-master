package main

import (
	"fmt"
	"strings"
)

func main() {
	s := ""
	fmt.Scan(&s)
	arr := strings.Split(s, "-")

	sum := int(arr[0][0]-'0') * 1

	for i := 0; i < len(arr[1]); i++ {
		sum += int(arr[1][i]-'0') * (i + 2)
	}

	for i := 0; i < len(arr[2]); i++ {
		sum += int(arr[2][i]-'0') * (i + 5)
	}

	m := sum % 11

	if m == 10 && arr[3][0] != 'X' {
		fmt.Printf("%s-%s-%s-%s", arr[0], arr[1], arr[2], "X")
		return
	}

	if m != 10 && int(arr[3][0]-'0') != m {
		fmt.Printf("%s-%s-%s-%d", arr[0], arr[1], arr[2], m)
		return
	}

	fmt.Println("Right")

}
