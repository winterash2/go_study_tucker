package main

import (
	"fmt"
)

// func main(){
// 	var char rune = '한'
// 	fmt.Printf("%T\n", char) // char의 타입 출력
// 	fmt.Println(char) // char값 출력
// 	fmt.Printf("%c\n", char) //문자 출력
// }

// func main() {
// 	str := "Hello 월드!"

// 	for i := 0; i < len(str); i++ {
// 		fmt.Printf("index: %d, 타입: %T, 값: %d, 문자값: %c \n", i, str[i], str[i], str[i])
// 	}
// }

// func main() {
// 	str := "Hello 월드!"
// 	strRune := []rune(str)

// 	for i := 0; i < len(strRune); i++ {
// 		fmt.Printf(" 타입: %T, 값: %d, 문자값: %c \n", strRune[i], strRune[i], strRune[i])
// 	}
// }

// func main() {
// 	str := "Hello 월드!"
// 	for _, v := range str {
// 		fmt.Printf("타입: %T, 값: %d, 문자: %c\n", v, v, v)
// 	}
// }

func main() {
	strHello := "Hello"
	str월드 := "월드"
	strHello월드 := strHello + " " + str월드
	fmt.Println(strHello월드)

	strHello += " " + str월드
	fmt.Println(strHello)
}
