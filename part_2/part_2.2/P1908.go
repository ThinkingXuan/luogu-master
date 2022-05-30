package main

import (
	"bufio"
	"fmt"
	"os"
)

//https://www.luogu.com.cn/problem/P1908
//逆序对

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
	fmt.Println(mergeSort(0, n-1, nums, make([]int, n)))
}

// 归并排序
func mergeSort(left, right int, nums []int, tmp []int) int {
	if left >= right {
		return 0
	}

	mid := (left + right) / 2
	res := mergeSort(left, mid, nums, tmp) + mergeSort(mid+1, right, nums, tmp)

	// 合并阶段
	i, j := left, mid+1
	for k := left; k <= right; k++ {
		tmp[k] = nums[k]
	}
	for k := left; k <= right; k++ {
		if i == mid+1 {
			nums[k] = tmp[j]
			j++
		} else if j == right+1 || tmp[i] <= tmp[j] {
			nums[k] = tmp[i]
			i++
		} else {
			nums[k] = tmp[j]
			j++
			res += mid - i + 1
		}
	}
	return res
}
