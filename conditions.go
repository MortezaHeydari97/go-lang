package main

import (
	"fmt"
	"strings"
	"math"
)

// `if` statement do not needs to parentheses but reqires the braces
func GtLt(operation, text string, count int) bool {
	if operation == "more" {
		if strings.Count(text, "") > count {
			return true
		}
		return false
	} else if operation == "less" {
		if strings.Count(text, "") < count {
			return true
		}
		return false
	}
	return false
}


/*
 * If with a short statement
 * The `if` statement like `for` can start with a short statement to excute before the condition.
 * The short declared statement can be used only in the `if` and `else` scopes
**/
func pow(x, y, limit float64) float64 {
	if v:= math.Pow(x, y); v < limit {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, limit)
	}

	return limit
}




func main() {
	var more = GtLt("more" , "Go Programming Language !", 10)
	var less = GtLt("less", "Go Programmin Language !", 10)
	fmt.Println("result =>", more)
	fmt.Println("result =>", less)
}