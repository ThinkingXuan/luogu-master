package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

//https://www.luogu.com.cn/problem/P1902
// 由题目，求最大值最小 所以用二分
// 解法： 答案满足单调，有界， 所以从答案入手进行二分，找到ans
// 对ans进行bfs查找，确定是最小值
const (
	MaxN = 1005
)

var dx = []int{1, 0, -1, 0}
var dy = []int{0, 1, 0, -1}

var p [MaxN][MaxN]int
var vis [MaxN][MaxN]bool
var n, m int

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
	var res int
	readInt(&n)
	readInt(&m)
	maxV, minV := math.MinInt32, math.MaxInt32

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			var tmp int
			readInt(&tmp)
			p[i][j] = tmp

			maxV = max(maxV, tmp)
			minV = min(minV, tmp)
		}
	}

	left, right := minV, maxV

	for left <= right {
		mid := (left + right) >> 1
		// vis 清0
		clear()
		if bfs(0, 0, mid) {
			right = mid - 1
			res = mid
		} else {
			left = mid + 1
		}
	}
	fmt.Println(res)
}

type pair struct {
	first  int
	second int
}

func bfs(x, y int, maxN int) bool {
	var queue []pair
	queue = append(queue, pair{x, y})
	vis[x][y] = true

	for len(queue) > 0 {
		xx := queue[0].first
		yy := queue[0].second
		queue = queue[1:]

		for i := 0; i < 4; i++ {
			nx := xx + dx[i]
			ny := yy + dy[i]

			if nx < 0 || nx >= n || ny < 0 || ny >= m || vis[nx][ny] || p[nx][ny] > maxN {
				continue //  不可行（越界、已访问、伤害过大）的直接跳过
			}

			vis[nx][ny] = true
			if nx == n-1 {
				return true //到了，返回true
			} else {
				queue = append(queue, pair{nx, ny})
			}
		}
	}
	return false // //没有搜到
}

func clear() {
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			vis[i][j] = false
		}
	}
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
