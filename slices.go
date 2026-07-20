package main

import (
	"fmt"
	"slices"
)

// Slices are important datatypes in Go, giving a more powerful
// interfaces to sequences than arrays and are dynamic too

func main() {

	// unline Arrays, Slices are typed only by the elements they contains
	// (not the number of length). An uninitilized slice equals to nil and
	// has length 0 ( zero ).
	var s []string
	fmt.Println("Uninit : ", s, s == nil, len(s) == 0)

	/* To create slices with non-zero length, use the builtin make. Here we
	make a slice of stringS of length 3 ( initially zero-values). By default a new
	slice's capacity is equal to its length; If we know the slice is going to grow
	ahead of time, it' possible to pass a capacity explicitly as an additional parameter
	to make.*/
	s = make([]string, 3)
	fmt.Println("Emp : ", s, ",len : ", len(s), ",cap : ", cap(s))

	// We can se nd get just like with arrays
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set : ", s)
	fmt.Println("get : ", s[2])

	// len returns the length of the slice as expected
	fmt.Println("len : ", len(s))

	/* In addition to these basic operation, slices support
	several more that make them richer than arrays, One is the
	builtin append, which returns a slice containing one or more
	new vaalues. Note that we need to accept a return value from
	append as we may get a new slice value.	*/

	fmt.Println("S before append : ", s)
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("s after append : ", s)

	// Slices can also be copy'd. Here we creaate an empty slice
	// c of the length as s and copy s into c.
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("Copy : ", c)

	/* Slices support "slice" operator with the syntax
	slice[low:high]. For example, this gets a slice of
	the element s[2], s[3], and s[4]. */
	l := s[2:5]
	fmt.Println("L : ", l)

	// This slices up to (but excluding) s[5]
	l = s[:5]
	fmt.Println("l : ", l)

	// And this slices up from (and including) s[2].
	l = s[2:]
	fmt.Println("l : ", l)

	// We can declare and initialize a variable for slice in a single
	// line as well.
	t := []string{"g", "h", "i", "j"}
	fmt.Println("t : ", t)

	// The slices package contains a number of useful utility functions
	// for slices
	t2 := []string{"g", "h", "i", "j"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	/* Slices can be composed into multi-dimensional data structures. The length
	of the inner slices can vary, unlike wih multi-dimensional arrays.*/

	twoD := make([][]int, 3)
	for i := range 3 {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := range innerLen {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d : ", twoD)
}
