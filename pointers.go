package main

import (
	"fmt"
)


func main() {
	/*
	 * Pointer defenition: A pointer is a variable which points a memory address.
	 * Here is an example to how create a pointer
	**/
	
	// var p *int
	// its defualt value is nil
	var a int = 10

	var x *int = &a

	// These two are equal
	fmt.Println("result => ", &a, x)

	// here is a shorthand:
	h := "morteza"
	m := &h

	/*
	 * The difference between a variable and pointer is that a variable stores the value at a memory address and a pointer points to a memory address
	 * You can change the variable value using a pointer like so:
	**/
	var b int = 100
	var y *int = &b

	*y = 15 // The 

	fmt.Println("result => ", b, *m) // its value changed
	

	/*
	 * You can also create a pointer using the `new` built-in function
	**/
	pa := new(int)

	fmt.Println("new built-in function ", *pa)
}