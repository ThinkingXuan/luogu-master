package main

import (
	"fmt"
)

//P2010 [NOIP2016 普及组] 回文日期

func main() {
	var data1, data2 int
	fmt.Scan(&data1, &data2)
	res := 0
	for i := data1; i <= data2; i++ {
		// 判断日期是否正确
		// 判断是否使回文
		if isData(i) && f(fmt.Sprintf("%d", i)) {
			res++
		}
	}
	fmt.Println(res)
}

func isData(date int) bool {
	//20111102
	year := date / 10000
	month := date / 100 % 100
	day := date % 100

	if month <= 0 || month > 12 {
		return false
	}
	if day <= 0 || day > 31 {
		return false
	}

	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		if day <= 31 {
			return true
		} else {
			return false
		}
	case 4, 6, 9, 11:
		if day <= 30 {
			return true
		} else {
			return false
		}
	case 2:
		if f2(year) {
			if day <= 29 {
				return true
			} else {
				return false
			}
		} else {
			if day <= 28 {
				return true
			} else {
				return false
			}
		}
	}
	return false
}

// 瑞年
func f2(year int) bool {
	if year%400 == 0 || (year%100 != 0 && year%4 == 0) {
		return true
	}
	return false
}

func f(date string) bool {
	n := len(date)
	for i := 0; i < n/2; i++ {
		if date[i] != date[n-i-1] {
			return false
		}
	}
	return true
}

// 方法2， 讨巧方法， 年份给的太少。
/*
// 没有考虑闰年，可以加以改进
// 遍历月和日，反推构成回文的年。
var (
    days = []int{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
)

func countPalindromicDate(start, end int) int {
    count := 0
    for month := 1; month <= 12; month++ {
        for day := 1; day <= days[month - 1]; day++ {
            year := (day % 10) * 1000 + (day / 10) * 100 + (month % 10) * 10 + month / 10
            date := year * 10000 + month * 100 + day
            if date >= start && date <= end {
                count++
            }
        }
    }
    return count
}

func main() {
    var start, end int
    fmt.Scanln(&start)
    fmt.Scanln(&end)
    fmt.Println(countPalindromicDate(start, end))
}
*/
