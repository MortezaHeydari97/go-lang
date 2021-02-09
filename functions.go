package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x + y)
	}
	
	fmt.Println("result => ", hypot(1.1, 2.3))
	fmt.Println("result => ", compute(hypot))
}