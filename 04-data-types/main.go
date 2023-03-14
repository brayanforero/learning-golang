package main

import "fmt"

func main() {

	// Data types
	// string, numeric, bool

	// numerics
	// ---------------------------------------------------
	// int8	8-bit signed integer
	// int16	16-bit signed integer
	// int32	32-bit signed integer
	// int64	64-bit signed integer

	// uint8	8-bit unsigned integer
	// uint16	16-bit unsigned integer
	// uint32	32-bit unsigned integer
	// uint64	64-bit unsigned integer

	// int	Both int and uint contain same size, either 32 or 64 bit.

	// uint	Both int and uint contain same size, either 32 or 64 bit.
	// rune	It is a synonym of int32 and also represent Unicode code points.

	// byte	It is a synonym of uint8.

	// uintptr	It is an unsigned integer type. Its width is not defined, but its can hold all the bits of a pointer value.

	// floats

	// float32	32-bit IEEE 754 floating-point number
	// float64	64-bit IEEE 754 floating-point number

	// Using 8-bit unsigned int
	var X uint8 = 225
	fmt.Println(X, X-3)

	// Using 16-bit signed int
	var Y int16 = 32767
	fmt.Println(Y+2, Y-2)

	a := 20.45
	b := 34.89

	// Subtraction of two
	// floating-point number
	c := b - a

	// Display the result
	fmt.Printf("Result is: %f", c)

	// Display the type of c variable
	fmt.Printf("\nThe type of c is : %T", c)

	// BOOELANS AND STRINGS

	// variables
	str1 := "LearnGO"
	str2 := "GOisAwesone"
	str3 := "LearnGO"
	result1 := str1 == str2
	result2 := str1 == str3

	// Display the result
	fmt.Println(result1)
	fmt.Println(result2)

	// Display the type of
	// result1 and result2
	fmt.Printf("The type of result1 is %T and "+
		"the type of result2 is %T",
		result1, result2)
}
