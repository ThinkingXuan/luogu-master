package main

import (
	"bufio"
	"fmt"
	"os"
)

var in = bufio.NewReader(os.Stdin)

func readInt(out *int) error {
	var ans, sign = 0, 1
	var readed = false
	c, err := in.ReadByte()
	for ; err == nil && (c < '0' || '9' < c); c, err = in.ReadByte() {
		if c == '-' {
			sign = -sign
		}
	}
	for ; err == nil && '0' <= c && c <= '9'; c, err = in.ReadByte() {
		ans = ans<<3 + ans<<1 + int(c-'0')
		readed = true
	}
	if readed {
		*out = ans * sign
		return nil
	}
	return err
}

func main() {
	var n, m, k int
	readInt(&n)
	readInt(&m)
	readInt(&k)
	// 初始化map集合
	numM := make([]map[int]struct{}, k)
	for i := range numM {
		numM[i] = make(map[int]struct{})
	}
	for i := 0; i < n; i++ {
		var tmp int
		for j := 0; j < m; j++ {
			readInt(&tmp)
			numM[tmp-1][j] = struct{}{}
		}
	}

	for i := 0; i < k; i++ {
		if i == 0 {
			fmt.Printf("%d", len(numM[i]))
		} else {
			fmt.Printf(" %d", len(numM[i]))
		}
	}

	//res := make([]int, k)
	//for _, mp := range numM {
	//	for key := range mp {
	//		res[key-1]++
	//	}
	//}
	//for _, v := range res {
	//	fmt.Print(v, " ")
	//}
}
