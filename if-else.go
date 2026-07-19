package main

import "fmt"

func main() {
	//basic example

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// if statement without else
	if 8%4 == 0 {
		fmt.Println("8 is even")
	}

	//logic operators like && and || re often used in the condition
	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("Either 7 or 8 is Odd")
	}

	/* A statement can precede conditions; any variables declared in this statement
	   are available in the current and all subsequent branches */
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has i digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

}
