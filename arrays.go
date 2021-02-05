package main 

import "fmt"


func main() {

	/*
	 * In go, type [n]T is an array of n values of type T
	 * This expression will declares a variable x as an array of 3 integers:
	 * var x [3]int
	 */
	
	/*
	 * An array's length is part of its type, so arrays cannot be resized. 
	 * This seems limiting, but don't worry; Go provides a convenient way of working with arrays.
	 */

	var fullName [2]string
	fullName[0] = "Morteza"
	fullName[1] = "Heydari"

	fmt.Println("Your full name is =>", fullName)


	primes := [4]int{10, 20, 30, 40}
	fmt.Println("primes arr is =>", primes)
}