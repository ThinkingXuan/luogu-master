package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

// 数组合并
func merge(A []player, B []player, res []player) {
	n := len(A)
	i, j, k := 0, 0, 0

	for i < n && j < n {
		if A[i].score > B[j].score || (A[i].score == B[j].score && A[i].number < B[j].number) {
			res[k] = A[i]
			k++
			i++
		} else {
			res[k] = B[j]
			k++
			j++
		}
	}

	for i < n {
		res[k] = A[i]
		k++
		i++
	}
	for j < n {
		res[k] = B[j]
		k++
		j++
	}
}

type player struct {
	score  int
	power  int
	number int
}

func main() {
	var n, r, q int
	//fmt.Scan(&n, &r, &q)
	readInt(&n)
	readInt(&r)
	readInt(&q)

	res := make([]player, 2*n)

	for i := 0; i < 2*n; i++ {
		readInt(&res[i].score)
		res[i].number = i + 1
	}

	for i := 0; i < 2*n; i++ {
		readInt(&res[i].power)
	}

	// 先自定义排序
	sort.Slice(res, func(i, j int) bool {
		if res[i].score == res[j].score {
			return res[i].number < res[j].number
		}
		return res[i].score > res[j].score
	})

	A := make([]player, n)
	B := make([]player, n)

	for i := 0; i < r; i++ {

		for j := 0; j < n; j++ {
			if res[2*j].power > res[2*j+1].power {
				res[2*j].score++
				A[j] = res[2*j]
				B[j] = res[2*j+1]
			} else {
				res[2*j+1].score++
				A[j] = res[2*j+1]
				B[j] = res[2*j]
			}
		}

		merge(A, B, res)
	}
	fmt.Println(res[q-1].number)
}
