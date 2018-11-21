package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	var b bool = true
	b = false
	fmt.Println(b)
	var f float32 = 231232131232132131231233243333333
	fmt.Println(f)
	var s string = "123321312"
	fmt.Println(s)
	abc("123")
	var d = "123"
	fmt.Println(d)
	e := "123333333333333333333"
	fmt.Println(e)
}

func abc(c string){
	fmt.Println(c)
}
