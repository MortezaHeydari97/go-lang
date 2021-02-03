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


	// We also can access the struct fields using pointers.
	v := Primary{10, 20}
	p := &v
	p.x = 1e3

	fmt.Println("cumbersome", *&p.x) // We can use struct field using a pointer but this is not good at all
	fmt.Println("new", p.x) // This is better to access struct field using pointer

	// Struct literals 
	v1 := Primary{7, 8} // x: 7, y: 8
	v2 := Primary{x: 5} // x: 5, y: 0
	v3 := Primary{} // x: 0, y: 0
	v4 := &Primary{4, 5} // x

	fmt.Println("a", v1)
	fmt.Println("b", v2)
	fmt.Println("c", v3)
	fmt.Println("d", *v4)
}