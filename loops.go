package main

import (
	"fmt"
)

func main() {
	var limit = 100
	var sum int16 = 0

	/*
	 * Unlike other languages like javascript, `for` statement do not need to any parantheses but requires the breaces
	 * The declared variables in `for` statement only are accessible inside its block
	**/
	for i := 0; i < limit; i++ {
		sum++
	}


	// The first and second parts of the `for` statement are optional
	summary := 1
	for ; summary < 100; {
		summary += summary
	}

	// And the `while` is like so:
	start := 0
	for start < 100 {
		start++
	}

	// A non-stop loop
	infinite := 0
	for {
		infinite++
	}


	fmt.Print(sum)
}