package main

import "fmt"

// Variadic functions can be called with any number of trailing arguments.
// For ex: fmt.Println is a common variadic function

// Here's a function that will take an arbitrary number of ints as arguments
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	// With in the function, the type of nums is equivalent to  []int. We can
	// call len(nums), iterate over it with range, etc.
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	// Variadic functions can be called in teh usual way with individual arguments
	sum(1, 2)
	sum(1, 2, 3)

	// If you already have multiple args in a slice, apply them to a variadic function
	// using func(slices..) like this.
	nums := []int{1, 2, 3, 4}
	sum(nums...)

}
