package main

import (
	"fmt"
)

func switch_boolean() {
	temp := 5

	switch true {
	case false:
		fmt.Println("조건에 맞는 문장이 없습니다")
	case temp < 10, temp > 30: // 한 번에 여러값 비교를 이용하여 or와 동일한 효과
		fmt.Println("너무 덥거나 너무 춥습니다.")
	case temp >= 10 && temp <= 30:
		fmt.Println("적당한 온도입니다.")
	default:
		fmt.Println("나올 수 없음")
	}
	fmt.Println("프로그램 종료")
}

func getMyAge() int {
	return 31
}
func initValue() {
	switch age := getMyAge(); age {
	case 31:
		fmt.Println("31살입니다.")
	}
}

type ColorType int

const (
	Red ColorType = iota
	Blue
	Green
	Yellow
)

func colorToString(color ColorType) string {
	switch color {
	case Red:
		return "Red"
	case Blue:
		return "Blue"
	case Green:
		return "Green"
	case Yellow:
		return "Yellow"
	default:
		return "Undefined"
	}
}

func main() {
	switch a := 1; a {
	case 1:
		fmt.Println("a == 1")
		fallthrough
	case 2:
		fmt.Println("a == 2")
	case 3:
		fmt.Println("a == 3")
	}
}
