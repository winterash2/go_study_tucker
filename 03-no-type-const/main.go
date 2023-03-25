package main

import "fmt"

const PI = 3.14
const FloatPI float64 = 3.14

func main() {
	var a int = PI * 100      // 3.14*100은 314로 정수이기 때문에 error 발생 X
	var b int = FloatPI * 100 // float * 100은 타입이 float이기 떄문에 int에 할당 못해서 에러 발생

	fmt.Println(a)
	fmt.Println(b)
}
