package main

import "fmt"

func main() {

	// () / * % + - -> Math operators
	// && || == -> Logic operators
	// ++ -- -> Post and pre increment
	// = -> to assigment a value in a variable

	a := 12
	b := 14
	c := a + b

	fmt.Printf("The result is: %d", c)

	age := 18

	isChild := age > 18

	fmt.Printf("\nIs the child: %v", isChild)

	number := 28

	number++

	fmt.Printf("\nThe number value is: &%d", number)

	// Do not do opertaions with variables that containt equals values

	var d int8 = 18

	e := 360
	// this operation is no valid, we must do a cating value
	// f := d + e

	f := int(d) + e
	fmt.Printf("\nThe value of variable 'F' is: %d", f)
}
