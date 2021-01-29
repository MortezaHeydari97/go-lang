package main
/* Here is all golang datatypes */


// String datatype
var name string = "Morteza"
var description string = `This is a test to write multiline string.
Author: morteza heydari`


// Boolean datatype
var isAlive bool = true



/*
 * Numeric datatypes
 * There are different types of numeric types
 * These types are different based on their storing bit size
 *
 * in8		- 8  bit - range => -128 to 127
 * int16	- 16 bit - range => math.Pow(-2, 15) to math.Pow(2, 15) -1
 * int32	- 32 bit - range => math.Pow(-2, 31) to math.Pow(2, 31) -1
 * int64	- 64 bit - range => math.Pow(-2, 63) to math.Pow(2, 63) -1
 * int 		- Based on the platform( 32 bits wide in 32-bit platforms and 64 bits wide in 64-bit platforms) 
 *
 *
 * unit8 	- 8  bit - range => 0 to 255
 * unit16 	- 16 bit - range => 0 to math.Pow(2, 16) - 1
 * unit32 	- 32 bit - range => 0 to math.Pow(2, 32) - 1
 * unit64 	- 64 bit - range => 0 to math.Pow(2, 64) - 1
 * unit 	- Based on the platform( 32 bits wide in 32-bit platforms and 64 bits wide in 64-bit platforms) 
 *
 * byte 	- alias for unit8
 * rune 	- alias for int32
 * 
 *
 * float32 	- occupies 32 bits in the memory
 * float64 	- occupies 64 bits in the memory
*/

var integer = 100 // default type is to int

var float = 200.2 // default type is float64



func main() { }