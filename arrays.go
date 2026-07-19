package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("emp:", a)

	//We can set a value at an index using the array[index] = value
	//syntax, and get a value with array[index]
	a[4] = 100
	fmt.Println("emp:", a)
	fmt.Println("get:", a[4])

	//The builtin len returns the length of an array
	fmt.Println("Length of Array :", len(a))

	//use this syntax to declare and initialize an array in one line.
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl :", b)

	//You can also have the compiler count the number of
	// elements for you with ...
	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl :", b)

	//If you specify the index with :, the elements in between
	// will be zeroed.
	c := [...]int{100, 3: 400, 500}
	fmt.Println("idx ", c)

	/* Array types are one-dimensional, but you can compose types to
	build multiple dimensional data structures.*/
	var twoD [2][3]int
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d : ", twoD)

	twoD = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("new 2d : ", twoD)

}
