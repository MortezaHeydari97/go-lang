package main

import (
	"fmt"
)

// A defer statement defers the execution of a function until the surrounding function returns.
func deferShow() {
	defer fmt.Println("World!")

	fmt.Println("Hello")
}

func main() {
	deferShow()


	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}