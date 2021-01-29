package main

import "fmt"

func main() {
	/*
	 * All types of operators in golang are like so:
	 * 
	 * 1. Arithmetic Operators
	 * 2. Assignment Operators
	 * 3. Comparison Operators
	 * 4. Logical Operators
	 * 5. Bitwise Operators
	*/
	
	
	// Arithmetic Operators Examples
	var (
		a	int8 = 1 + 1	// Addition
		b 	int8 = 5 - 1	// Substraction
		c 	int8 = 2 * 3	// Multiplication
		d	int8 = 10 / 2 	// Division
		e	int8 = 10 % 2	// Modulus
	)
	c++	// Increment
	c--	// Decrement
	
	
	// Assignment Operators
	a += b // Add and assign
	a -= b // Subtract and assign
	a *= b // Multiply and assign
	a /= b // Divide and assign quotient
	a %= b // Divide and assign modulus
	
	
	// Comparison Operators
	var (
		f = a == b 	// Equal
		g = a != b 	// Not equal
		h = a < b 	// Less than
		i = a <= b 	// Less than or equal to
		j = a > b 	// Greater than
		k = a >= b  // Greater than or equal to
	)
	
	
	// Logical Operators
	var l = a < b && d > e
	var m = a < b || d > e
	
	
	// Bitwise Operators
	var (
		n = a &	 b // AND
		o = a |	 b // OR
		p = a ^	 b // XOR
		q = a << b // Zero fill left shift	
		r = a >> b // Signed right shift
	)
	
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(l)
	fmt.Println(m)
	fmt.Println(n)
	fmt.Println(o)
	fmt.Println(p)
	fmt.Println(q)
	fmt.Println(r)
}