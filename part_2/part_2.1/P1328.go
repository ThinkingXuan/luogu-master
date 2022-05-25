package main

import (
	"fmt"
)

func main() {
	var n, na, nb int
	fmt.Scan(&n, &na, &nb)

	numa := make([]int, na)
	numb := make([]int, nb)

	for i := 0; i < na; i++ {
		fmt.Scan(&numa[i])
	}

	for j := 0; j < nb; j++ {
		fmt.Scan(&numb[j])
	}
	resA, resB := 0, 0
	cnt := 0
	ai, bi := 0, 0
	for cnt < n {
		res := judge(numa[ai], numb[bi])
		if res == 1 {
			resA++
		} else if res == 2 {
			resB++
		}
		ai = (ai + 1) % na
		bi = (bi + 1) % nb
		cnt++
	}
	fmt.Println(resA, resB)

}

// 	//0 表示“剪刀”，1表示“石头”，2 表示“布”，3 表示“蜥蜴人”，4表示“斯波克”。
// 0 平局，1 a胜b ,  2 a负b
func judge(a, b int) int {
	switch a {
	case 0:
		switch b {
		case 0:
			return 0
		case 1:
			return 2
		case 2:
			return 1
		case 3:
			return 1
		case 4:
			return 2
		}
	case 1:
		switch b {
		case 0:
			return 1
		case 1:
			return 0
		case 2:
			return 2
		case 3:
			return 1
		case 4:
			return 2
		}
	case 2:
		switch b {
		case 0:
			return 2
		case 1:
			return 1
		case 2:
			return 0
		case 3:
			return 2
		case 4:
			return 1
		}
	case 3:
		switch b {
		case 0:
			return 2
		case 1:
			return 2
		case 2:
			return 1
		case 3:
			return 0
		case 4:
			return 1
		}
	case 4:
		switch b {
		case 0:
			return 1
		case 1:
			return 1
		case 2:
			return 2
		case 3:
			return 2
		case 4:
			return 0
		}
	}
	return -1
}
