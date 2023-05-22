package main

import (
	"fmt"
	"unsafe"
)

// type Person struct {
// 	Name string
// 	Age  int
// }

// type Student struct {
// 	Class      int
// 	No         int
// 	Score      float64
// 	PersonInfo Person
// }

// func test1() {
// 	fmt.Println("struct 사용 예시")

// 	var gildong Student
// 	gildong.Class = 2
// 	fmt.Println(gildong) // {2 0 0 { }}

// 	var choelsu Student = Student{1, 2, 3, Person{}}
// 	fmt.Println(choelsu) // {1 2 3 { }}

// 	var younghee = Student{Class: 3}
// 	fmt.Println(younghee) // {3 0 0 { }}

// 	hwarangPerson := Person{"송화랑", 12}
// 	hwarangStudent := Student{1, 2, 3, hwarangPerson}
// 	fmt.Println(hwarangStudent) // {1 2 3 {송화랑 12}}
// }

type Person struct {
	Name string
	Age  int
}

type Student struct {
	Class      int
	No         int
	Score      float64
	Person
}

func main() {
	gildong := Student{1, 2, 3, Person{"홍길동", 2}}
	fmt.Println(gildong.Name)
	fmt.Println(unsafe.Sizeof(gildong))
	gildong2 := gildong
	gildong2.Name = "홍길동2"
	fmt.Println(gildong)
	fmt.Println(gildong2)
}

// type Person struct {
// 	Name string
// 	Age  int
// }

// type Onlyname struct {
// 	Name string
// }

// type Student struct {
// 	Class      int
// 	No         int
// 	Score      float64
// 	Person
// 	Onlyname
// }

// func main() {
// 	gildong := Student{1, 2, 3, Person{"홍길동", 2}, Onlyname{"김철수"}}
// 	fmt.Println(gildong.Name)
// }

