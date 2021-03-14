package main

import (
	"fmt"
	"time"

	"github.com/aclements/go-moremath/stats"
)

func trapezoid(a, b, h float64) float64 {
	return ((a + b) * h) / 2
}

const SECTION int = 65536 // 分割数

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

// 積分範囲をn等分してできる台形の面積の和を足して数値積分を行う
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

// 標準正規分布表を作る
func test3() {
	norm := stats.NormalDist{0, 1} // 平均0, 分散1の正規分布
	z1 := []float64{0.0, 0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9, 1.0, 1.1, 1.2, 1.3, 1.4, 1.5, 1.6, 1.7, 1.8, 1.9, 2.0, 2.1, 2.2, 2.3, 2.4, 2.5, 2.6, 2.7, 2.8, 2.9, 3.0}
	z2 := []float64{0.00, 0.01, 0.02, 0.03, 0.04, 0.05, 0.06, 0.07, 0.08, 0.09}

	fmt.Println("  z  |  0.00   0.01   0.02   0.03   0.04   0.05   0.06   0.07   0.08   0.09 ")
	fmt.Println("----------------------------------------------------------------------------")
	for _, i := range z1 {
		fmt.Printf(" %.1f | ", i)
		for _, j := range z2 {
			fmt.Printf("%.4f ", integrate2(norm.PDF, 0, (i+j)))
		}
		fmt.Printf("\n")
	}
}

// 数値積分の高速化
func integrate3(f func(float64) float64, a, b float64) float64 {
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

func stopwatch(f func()) {
	start := time.Now()

	for i := 0; i < 10; i++ {
		f()
	}

	end := time.Now().Sub(start)
	fmt.Printf("Time : ")
	fmt.Println(end / 10)
}

func main() {
	stopwatch(test3)
}
