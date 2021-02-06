package main

import (
	"fmt"
	"time"
)

/*
 * A struct is a collection of fields.
 * Here is an example of how to create a struct
 */
type Primary struct {
	x int
	y int
}

type getAge func(int) int

type getFullName func(string, string) string

type User struct {
	firstName string
	lastName string
	age getAge
	fullName getFullName
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


	// A more complex and real example to working with strcuts
	userInfo := User{
		firstName: "Morteza",
		lastName: "Heydari",
		fullName: func(fName, lName string) string {
			full := fmt.Sprintf("%v %v", fName, lName)
			return full
		},
		age: func(year int) int {
			t := time.Now()
			yr := t.Year()
			return yr - year
		},
	}
	fmt.Printf("My name is %v and I'm %v years old.", userInfo.fullName(userInfo.firstName, userInfo.lastName), userInfo.age(1997))
}