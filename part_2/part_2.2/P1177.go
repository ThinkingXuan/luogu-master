package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

//https://www.luogu.com.cn/problem/P1177

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
	var n int
	readInt(&n)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		readInt(&nums[i])
	}
	quickSort(nums, 0, n-1)
	for i := 0; i < n; i++ {
		fmt.Print(nums[i], " ")
	}
}

func quickSort(nums []int, l, r int) {
	left, right := l, r
	if left >= right {
		return
	}

	// 随机选择一个下标
	randV := left + rand.Intn(right-left+1)
	nums[left], nums[randV] = nums[randV], nums[left]

	pivot := nums[left]

	for left < right {
		for left < right && nums[right] >= pivot {
			right--
		}
		nums[left] = nums[right]
		for left < right && nums[left] <= pivot {
			left++
		}
		nums[right] = nums[left]
	}

	nums[left] = pivot

	quickSort(nums, l, left-1)
	quickSort(nums, left+1, r)
}
