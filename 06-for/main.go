package main

import "fmt"

func main() {
OuterFor:
	for {
		for {
			break OuterFor
		}
	}
	fmt.Println("레이블을 이용해 중첩 for문 한번에 탈출")
}
