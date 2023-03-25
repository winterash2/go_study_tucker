package main

import (
	"fmt"
	"math"
	"math/big"
)

// 단순히 실수 비교를 한 경우
func check_error() {
	var a float64 = 0.1
	var b float64 = 0.2
	var c float64 = 0.3
	fmt.Println(a+b == c)
	fmt.Println(a + b)
}

// Nextafer를 이용하여 마지막 1비트만큼의 차이로 실수의 비교를 한 경우
func equal(a, b float64) bool {
	return math.Nextafter(a, b) == b
}

func diff_float_last_1_bit() {
	var a float64 = 0.1
	var b float64 = 0.2
	var c float64 = 0.3

	fmt.Println(equal(a+b, c))
}

func main() {
	a, _ := new(big.Float).SetString("0.1")
	b, _ := new(big.Float).SetString("0.2")
	c, _ := new(big.Float).SetString("0.3")

	d := new(big.Float).Add(a, b)
	fmt.Println(a, b, c, d) // -> 0.1 0.2 0.3 0.3
	fmt.Println(c.Cmp(d))   // -> 0(equal)
}
