package main

import (
	"fmt"
	"math"
)

const toggle = true

func main() {
	fmt.Println("toggle is set  to : ", toggle)

	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
