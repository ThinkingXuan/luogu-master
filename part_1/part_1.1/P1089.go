package main

import "fmt"

func main() {
	var num [12]int
	for i := 0; i < len(num); i++ {
		fmt.Scan(&num[i])
	}
	mom, my := 0, 0
	for i := 0; i < len(num); i++ {
		// 妈妈月初给300
		my += 300
		if num[i] > my { // 超支
			fmt.Println(-(i + 1))
			return
		} else { // 没有超支
			// 把所有的钱拿出来，剩余的钱
			remain := my - num[i]

			// 存入mom (整数)
			mom += (remain / 100) * 100
			// 保留在自己手上
			my = remain - (remain/100)*100
		}
	}
	fmt.Println(float64(mom)*1.2 + float64(my))
}
