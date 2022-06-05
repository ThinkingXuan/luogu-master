package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// 从答案二分+前缀和优化
const maxN = 200010

var w, v, l, r [maxN]int
var preN, preV [maxN]int64
var n, m, s int
var y, sum int64

func check(mid int) bool {
	y = 0
	sum = 0
	clear()

	for i := 1; i <= n; i++ {
		if w[i] >= mid {
			preN[i] = preN[i-1] + 1
			preV[i] = preV[i-1] + int64(v[i])
		} else {
			preN[i] = preN[i-1]
			preV[i] = preV[i-1]
		}
	}
	for i := 1; i <= m; i++ {
		y += (preN[r[i]] - preN[l[i]-1]) * (preV[r[i]] - preV[l[i]-1])
	}

	sum = abs64(y - int64(s))
	if y > int64(s) {
		return true
	}
	return false
}

func clear() {
	for i := 0; i <= n; i++ {
		preN[i] = 0
		preV[i] = 0
	}
}

func main() {
	minV, maxV := math.MaxInt32, math.MinInt32
	var ans int64 = math.MaxInt64

	fmt.Scan(&n, &m, &s)
	for i := 1; i <= n; i++ {
		readInt(&w[i])
		readInt(&v[i])
		minV = min(minV, w[i])
		maxV = max(maxV, w[i])
	}

	for i := 1; i <= m; i++ {
		readInt(&l[i])
		readInt(&r[i])
	}

	left, right := minV-1, maxV+2

	for left <= right {
		mid := (left + right) >> 1
		if check(mid) {
			left = mid + 1
		} else {
			right = mid - 1
		}
		if sum < ans {
			ans = sum
		}
	}

	fmt.Println(ans)

}

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

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func abs64(x int64) int64 {
	if x > 0 {
		return x
	}
	return -x
}
