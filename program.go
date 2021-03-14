package main

import (
	"fmt"
	"runtime"

	"github.com/aclements/go-moremath/stats"
)

func trapezoid(a, b, h float64) float64 {
	return ((a + b) * h) / 2
}

const SECTION int = 65536 // 分割数

// 数値積分の高速化
func integrate(f func(float64) float64, a, b float64) float64 {
	// 積分区間の分割
	if SECTION <= 1 {
		return trapezoid(f(a), f(b), (b - a))
	}

	section := make([]float64, SECTION)
	for i := 0; i < SECTION; i++ {
		section[i] = a + (((a + b) / float64(SECTION)) * float64(i+1))
	}

	cpu := runtime.NumCPU() // CPUの数
	runtime.GOMAXPROCS(cpu)

	ch := make(chan float64, cpu)
	for i := 0; i < cpu; i++ {
		// 処理の分割(範囲を8つに分けてCPUに割り振る)
		go func(i int) {
			var psum float64 = 0.0 // 部分和
			for j := i; j < SECTION-1; j += 8 {
				psum += trapezoid(f(section[j]), f(section[j+1]), (section[j+1] - section[j]))
			}
			ch <- psum
		}(i)
	}

	var sum float64 = 0.0
	for i := 0; i < cpu; i++ {
		sum += <-ch
	}

	return sum
}

func test() {
	norm := stats.NormalDist{0, 1} // 平均0, 分散1の正規分布
	z1 := []float64{0.0, 0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9, 1.0, 1.1, 1.2, 1.3, 1.4, 1.5, 1.6, 1.7, 1.8, 1.9, 2.0, 2.1, 2.2, 2.3, 2.4, 2.5, 2.6, 2.7, 2.8, 2.9, 3.0}
	z2 := []float64{0.00, 0.01, 0.02, 0.03, 0.04, 0.05, 0.06, 0.07, 0.08, 0.09}

	fmt.Println("  z  |  0.00   0.01   0.02   0.03   0.04   0.05   0.06   0.07   0.08   0.09 ")
	fmt.Println("----------------------------------------------------------------------------")
	for _, i := range z1 {
		fmt.Printf(" %.1f | ", i)
		for _, j := range z2 {
			fmt.Printf("%.4f ", integrate(norm.PDF, 0, (i+j)))
		}
		fmt.Printf("\n")
	}
}

func main() {
	test()
}
