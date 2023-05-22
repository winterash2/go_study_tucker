package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup // 패키지 전역으로 WaitGroup 객체를 하나 생성

func SumAtoB(a, b int) {
	sum := 0
	for i := a; i <= b; i++ {
		sum += i
	}
	fmt.Printf("%d부터 %d까지 합계는 %d입니다\n", a, b, sum)
	wg.Done()
}

func main() {
	wg.Add(10)                // 10개의 작업을 생성
	for i := 0; i < 10; i++ { // 10번 반복하면서 고루틴 생성
		go SumAtoB(1, 10000000) // 고루틴이 끝날때마다 Done을 호출
	}

	wg.Wait() // 앞서 만들었던 10개가 다 끝날때까지 대기
}
