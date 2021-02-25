package main

import "fmt"

/**
* Go by Example: Variables
*/
func main() {

	var a = "initial"
	fmt.Println(a)

	// var declares 1 or more variables.
	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	// the zero value for an int is 0.
	var e int
	fmt.Println(e)

	// The := syntax is shorthand for declaring and initializing a variable
	f := "apple"
	fmt.Println(f)

	
}