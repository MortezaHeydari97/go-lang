package main

import "fmt"

func main() {
	/*
	 * The range form of the for loop iterates over a slice or map.
     * When ranging over a slice, two values are returned for each iteration.
	 * The first is the index, and the second is a copy of the element at that index.
	 */

	// var pow = []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}
	
	// for i, v := range pow {
	// 	fmt.Printf("2**%d = %v\n", i, v)
	// }
	
	
	slc := make([]int, 10)
	fmt.Println("test", slc)


	// You can skip the index or value by assigning an underscore(_)
	for index := range slc {
		slc[index] = 1 << uint(index)
	}

	fmt.Println("now slc variable is => ", slc)
}