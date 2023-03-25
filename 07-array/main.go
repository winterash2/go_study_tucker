package main

import (
	"fmt"
)

func example1() {
	var arr [10]int = [10]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i])
	}
}

func example2() {
	var arr = [...]int{1, 2, 3}
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i])
	}
}

func example3() {
	arr := [...]int{1, 2, 3}
	for idx, value := range arr {
		fmt.Println("idx=", idx, ", value:", value)
	}
}

// func error(){
// 	a := [...]int{1, 2, 3}
// 	b := [5]int{}
// 	b = a
// }

func main() {
	a := [2][5]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
	}
	fmt.Println(a)
}
