package main

import (
	"fmt"
	"time"

	"github.com/aclements/go-moremath/stats"
)

func trapezoid(a, b, h float64) float64 {
	return ((a + b) * h) / 2
}

// 積分範囲a~bを高さとする台形の面積を求めることで数値積分を行う
func integrate1(f func(float64) float64, a, b float64) float64 {
	return trapezoid(f(a), f(b), (b - a))
}

func test1() {
	f := func(x float64) float64 {
		return x * x * x * x
	}
	fmt.Println(integrate1(f, 0, 5)) // 理論値 : 625
}

// 積分範囲を32等分してできる台形の面積の和を足して数値積分を行う
const SECTION int = 32 // 分割数

func integrate2(f func(float64) float64, a, b float64) float64 {
	// 積分区間の分割
	if SECTION <= 1 {
		return integrate1(f, a, b)
	}

	section := make([]float64, SECTION)
	for i := 0; i < SECTION; i++ {
		section[i] = a + (((a + b) / float64(SECTION)) * float64(i+1))
	}

	var sum float64 = 0
	for i := 0; i < SECTION-1; i++ {
		sum += trapezoid(f(section[i]), f(section[i+1]), (section[i+1] - section[i]))
	}
	return sum
}

func test2() {
	f := func(x float64) float64 {
		return x * x * x * x
	}
	fmt.Println(integrate2(f, 0, 5)) // 理論値 : 625
}

// t分布の両側α点を求める
func test3() {
	v := stats.TDist{1.0} // 自由度1.0
	fmt.Println(v)
}

func main() {
	start := time.Now()

	// ここに書く
	test3()

	end := time.Now().Sub(start)
	fmt.Printf("Time : ")
	fmt.Println(end)
}
