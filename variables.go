package main

import (
	"fmt"
)

func main() {

	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b + c)

	d := true
	fmt.Println("The value of d is : ", d)

	d = false
	fmt.Println("The value of d is updated to ", d)

	var e int
	fmt.Println("e : ", e)

	f := "apple"
	fmt.Println("f : ", f)
}
