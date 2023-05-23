package main

import (
	"fmt"
	"strconv"
)

//--------Person--------

type Person struct {
	Name string
	Age  int
}

func (p *Person) String() string {
	return "이름: " + p.Name + ", 나이: " + strconv.Itoa(p.Age)
}

//--------Car--------

type Car struct {
	Number string
}

func (c *Car) String() string {
	return "차량 번호: " + c.Number
}

// --------Stringer 인터페이스--------
type Stringer interface {
	String() string
}

func main1() {
	var stringer1 Stringer
	person1 := Person{"홍길동", 18}
	car1 := Car{"12가1234"}

	stringer1 = &person1
	fmt.Println(stringer1.String())
	stringer1 = &car1
	fmt.Println(stringer1.String())
}

func PrintVal(val interface{}) {
	switch t := val.(type) {
	case int:
		fmt.Printf("int: %d", t)
	case string:
		fmt.Printf("string: %s", t)
	}
}

func main() {
	valInt := 10
	valString := "Hello"

	PrintVal(valInt)
	PrintVal(valString)

	var interface1 interface{}
	interface1 = valInt
	valInt2 := interface1.(int)
	fmt.Println(valInt2)
}

i do it all the day
trick me really nicely