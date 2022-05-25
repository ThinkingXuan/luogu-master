package main

//import (
//	"bufio"
//	"fmt"
//	"os"
//)
//
//var in = bufio.NewReader(os.Stdin)
//
//func readInt(out *int) error {
//	var ans, sign = 0, 1
//	var readed = false
//	c, err := in.ReadByte()
//	for ; err == nil && (c < '0' || '9' < c); c, err = in.ReadByte() {
//		if c == '-' {
//			sign = -sign
//		}
//	}
//	for ; err == nil && '0' <= c && c <= '9'; c, err = in.ReadByte() {
//		ans = ans<<3 + ans<<1 + int(c-'0')
//		readed = true
//	}
//	if readed {
//		*out = ans * sign
//		return nil
//	}
//	return err
//}
//
//// readInt() 相比 fmt.Scan() 能提高很大的效率，数多的话，能提高10倍
//func main() {
//	var n int
//	readInt(&n)
//	num := make([][]int, n)
//	for i := 0; i < n; i++ {
//		var a, b, g, k int
//		readInt(&a)
//		readInt(&b)
//		readInt(&g)
//		readInt(&k)
//		num[i] = []int{a, b, g, k}
//	}
//
//	var x, y int
//	readInt(&x)
//	readInt(&y)
//
//	// 返回是最上面的编号
//	res := -1
//	for i := 0; i < n; i++ {
//		if x >= num[i][0] && x <= num[i][0]+num[i][2] && y >= num[i][1] && y <= num[i][1]+num[i][3] {
//			res = i + 1
//		}
//	}
//
//	fmt.Println(res)
//}
