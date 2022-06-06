package main

import (
	"bufio"
	"fmt"
	"os"
)

type lazyST []struct {
	l, r int
	todo int64
	sum  int64
}

func (lazyST) op(a, b int64) int64 {
	return a + b // % mod
}

func (t lazyST) maintain(o int) {
	lo, ro := t[o<<1], t[o<<1|1]
	t[o].sum = t.op(lo.sum, ro.sum)
}

func (t lazyST) build(a []int64, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t[o].sum = a[l-1]
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t lazyST) do(o int, add int64) {
	to := &t[o]
	to.todo += add                     // % mod
	to.sum += int64(to.r-to.l+1) * add // % mod
}

func (t lazyST) spread(o int) {
	if add := t[o].todo; add != 0 {
		t.do(o<<1, add)
		t.do(o<<1|1, add)
		t[o].todo = 0
	}
}

// 如果维护的数据（或者判断条件）具有单调性，我们就可以在线段树上二分
// 未找到时返回 n+1
// o=1  [l,r] 1<=l<=r<=n
// https://codeforces.com/problemset/problem/1179/C
func (t lazyST) binarySearch(o, l, r int, val int64) int {
	if t[o].l == t[o].r {
		if val <= t[o].sum {
			return t[o].l
		}
		return t[o].l + 1
	}
	t.spread(o)
	// 注意判断对象是当前节点还是子节点
	if val <= t[o<<1].sum {
		return t.binarySearch(o<<1, l, r, val)
	}
	return t.binarySearch(o<<1|1, l, r, val)
}

// o=1  [l,r] 1<=l<=r<=n
func (t lazyST) update(o, l, r int, add int64) {
	if l <= t[o].l && t[o].r <= r {
		t.do(o, add)
		return
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		t.update(o<<1, l, r, add)
	}
	if m < r {
		t.update(o<<1|1, l, r, add)
	}
	t.maintain(o)
}

// o=1  [l,r] 1<=l<=r<=n
func (t lazyST) query(o, l, r int) int64 {
	if l <= t[o].l && t[o].r <= r {
		return t[o].sum
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if m < l {
		return t.query(o<<1|1, l, r)
	}
	vl := t.query(o<<1, l, r)
	vr := t.query(o<<1|1, l, r)
	return t.op(vl, vr)
}

func (t lazyST) queryAll() int64 { return t[1].sum }

// a 从 0 开始
func newLazySegmentTree(a []int64) lazyST {
	t := make(lazyST, 4*len(a))
	t.build(a, 1, 1, len(a))
	return t
}

func main() {
	var n, m, sign int
	readInt(&n)
	readInt(&m)
	nums := make([]int64, n)
	for i := 0; i < n; i++ {
		readInt64(&nums[i])
	}

	seg := newLazySegmentTree(nums)

	for i := 0; i < m; i++ {
		readInt(&sign)
		if sign == 1 {
			var x, y int
			var k int64
			readInt(&x)
			readInt(&y)
			readInt64(&k)
			seg.update(1, x, y, k)
		} else {
			var x, y int
			readInt(&x)
			readInt(&y)
			fmt.Println(seg.query(1, x, y))
		}
	}

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

func readInt64(out *int64) error {
	var ans, sign int64 = 0, 1
	var readed = false
	c, err := in.ReadByte()
	for ; err == nil && (c < '0' || '9' < c); c, err = in.ReadByte() {
		if c == '-' {
			sign = -sign
		}
	}
	for ; err == nil && '0' <= c && c <= '9'; c, err = in.ReadByte() {
		ans = ans<<3 + ans<<1 + int64(c-'0')
		readed = true
	}
	if readed {
		*out = ans * sign
		return nil
	}
	return err
}
