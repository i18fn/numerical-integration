package main

import (
	"fmt"
	"time"
)

func Trapezoid(a, b, h float64) float64 {
	return ((a + b) * h) / 2
}

// 積分範囲a~bを高さとする台形の面積を求めることで数値積分を行う
func integrate1(f func(float64) float64, a, b float64) float64 {
	return ((f(a) + f(b)) * (b - a)) / 2
}

func test1() {
	f := func(x float64) float64 {
		return x * x
	}
	fmt.Println(integrate1(f, 2, 3)) // 理論値 : 7.66666666666
}

func main() {
	start := time.Now()

	// ここに書く
	test1()

	end := time.Now().Sub(start)
	fmt.Printf("Time : ")
	fmt.Println(end)
}
