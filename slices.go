package main

import (
	"fmt"
	"strings"
)

func main() {
	/*
	 * Slices are the tool you can create a dynamic sized array
	 * The type of []T is a slice with elements of type T
	 * A slice is formed by specifying two indices, a low and high bound, separated by a colon:
	 * slice[low : high]
	 */

	var arr = [5]string{"vue.js", "react.js", "angular2+", "go-gin", "revel"}
	var slc []string = arr[0:3]

	fmt.Println("slice", arr, slc)


	// Slice literals
	x := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(x)

	y := []bool{true, false, true, true, false, true}
	fmt.Println(y)

	// Here you can create a strcut and type it's name here instead of creating an inline struct
	z := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(z)
 
	 
	// Slice defaults
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[:]
	fmt.Println(s)
	
	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
 
	
	/*
	 * Slice length and capacity
	 * 
	 * The slice has both a length and a capacity
	 * The length of a slice is the number of elements it contains.
	 * The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
	 * 
	 * You can get the length and slice of a slice by `len` and `cap`
	 */
	n := []int{2, 3, 5, 7, 11, 13}
	fmt.Printf("length and capacity => %v and %v \n", len(n), cap(n))
	
	n = n[:0]
	fmt.Printf("length and capacity => %v and %v \n", len(n), cap(n))

	// Now if you write [x:] will get an error and you must to recovery the data like [:x] and now ypu can execute the [x:]
	n = n[:4]
	fmt.Printf("length and capacity => %v and %v \n", len(n), cap(n))

	n = n[1:]
	fmt.Printf("length and capacity => %v and %v \n", len(n), cap(n))
	

	/*
	 * Nil slice
	 * The zero value of a slice is nil
	 * The length and capacity of a nil slice is 0 
	 */
	
	var zr []int
	fmt.Printf("Now the result of zr == nil is %v, its value is %v, its length is %v and its capacity is %v\n", zr == nil, zr, len(zr), cap(zr))


	/*
	 * Create a slice with make
	 * The `make` is a one of go lang built-in functions that creates a slice and this is how to you can create a dynamiced-size array
	 */

	a := make([]int, 5)
	fmt.Println(a, len(a), cap(a))

	b := make([]int, 0, 3)
	fmt.Println(b, len(b), cap(b))

	
	
	/*
	 * Slices can contain any type, including other slices
	 * Let's create a tok-tac-toe board:
	 */
	
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
	

	/*
	 * Appending to a slice
	 * One of another go lang built-in functions is append.
	 * Its usage is append new elements into a slice.
	 * It takes some arguments that are like so:
	 * append(s []T, vs ...T) []T
	 * The first parameter is a slice of type T and the rest parameters are the new values you want to append them into the slice 
	 */

	var sl []int
	fmt.Println(sl, len(sl), cap(sl))

	sl = append(sl, 0)
	fmt.Println(sl, len(sl), cap(sl))

	sl = append(sl, 1)
	fmt.Println(sl, len(sl), cap(sl))

	sl = append(sl, 4, 5)
	fmt.Println(sl, len(sl), cap(sl))
}