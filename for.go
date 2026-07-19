package main

import "fmt"

func main() {
	i := 1

	// Basic for condiion
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	// Classic initial/condition/after for loop
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	// For loop without a condition keeps running until to break out of loop or return
	for {
		fmt.Println("loop")
		break
	}

	//Another approach of " Do this N times" iteration with rannge
	for i := range 3 {
		fmt.Println(i)
	}

	//For with continue to next iteration
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
