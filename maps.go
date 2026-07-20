package main

import (
	"fmt"
	"maps"
)

func main() {

	// To create an empty map, use the builtin make : make(map[key-type]val-type)
	m := make(map[string]int)

	// set key/value paris using typical name[key] = val syntax
	m["k1"] = 7
	m["k2"] = 13

	// printing a map with e.g.fmt.Printtln will show all of its key/value pairs.
	fmt.Println("map : ", m)

	// Get a value for key with name[key]
	v1 := m["k1"]
	fmt.Println("v1 : ", v1)

	// If the key doesn't exist, the zero value of the value type is returned.
	v3 := m["k3"]
	fmt.Println("v3 : ", v3)

	// The builtin len returns the number of key/value pairs when called on
	// a map
	fmt.Println("value of m : ", m)
	delete(m, "k2")
	fmt.Println("value of m after delete : ", m)

	// To remove all key/value paris from a map, use the clear builtin
	fmt.Println("value of m : ", m)
	clear(m)
	fmt.Println("value of m after clear : ", m)

	/* The optional second return value when getting a value from a map
	indicates if the key was present in the map. This can be used to
	disambiguate between missing keys an keys with zero value like 0 or "".
	Here we didn't need the value itself, so we ignored it with the blank
	identifier _. */
	fmt.Println("Data in m : ", m)
	_, prs := m["k2"]
	fmt.Println("prs : ", prs)

	// You can also declare and initilize a new map in the same line with
	// this syntax
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map : ", n)

	// The maps package contains a number of useful utility functions for maps.
	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}
