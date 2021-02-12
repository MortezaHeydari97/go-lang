package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}


// closure
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x + y)
	}
	
	fmt.Println("result => ", hypot(1.1, 2.3))
	fmt.Println("result => ", compute(hypot))

	pos, neg := adder(), adder()
	
	for i := 0 ;i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}
}