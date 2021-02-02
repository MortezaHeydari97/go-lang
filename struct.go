package main

import "fmt"

/*
 * A struct is a collection of fields.
 * Here is an example of how to create a struct
 */
type Primary struct {
	x int
	y int
}

func main() {
	var a = Primary{1, 2}

	// You can access a struct field using a dot like so:
	a.x = 5
	fmt.Println("result =>", a)
}