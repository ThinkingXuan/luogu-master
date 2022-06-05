package main

import "math/rand"

//https://leetcode.cn/problems/generate-random-point-in-a-circle/
// 给定圆的半径和圆心的位置，实现函数 randPoint ，在圆中产生均匀随机点。

type Solution struct {
	radius   float64
	x_center float64
	y_center float64
}

func Constructor(radius float64, x_center float64, y_center float64) Solution {
	return Solution{radius: radius, x_center: x_center, y_center: y_center}
}

// RandPoint 拒绝采样（循环的形式，直到满足条件）
func (this *Solution) RandPoint() []float64 {
	for {
		v_x := rand.Float64()*2 - 1
		v_y := rand.Float64()*2 - 1

		if v_x*v_x+v_y*v_y < 1 {
			return []float64{this.x_center + v_x*this.radius, this.y_center + v_y*this.radius}
		}
	}

}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(radius, x_center, y_center);
 * param_1 := obj.RandPoint();
 */
