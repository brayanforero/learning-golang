// This package talks about the pointers and what they are
package main

import (
	"fmt"
)

func main() {

	// What are they pointers
	// Pointers are a memory address that save value and reference of a previous declared variable.

	// How to use
	// 1 otherVariable := "This is the value of this variable"
	// 2.var myVar *datatype
	// 3.myVar = &otherVariable

	// & operator: &otherVariable
	//return -> place in memory like this: 0xc00009c020
	// * operator: *myVar
	//return -> value saved in the reference of variable: return the value of "otherVariable"

	fruit := "This a string"
	fmt.Printf("Type: %T, Value: %v, Address: %v", fruit, fruit, &fruit)

	var refFruit *string
	refFruit = &fruit

	fmt.Printf("\n\nType: %T, Value: %v, Address: %v", refFruit, *refFruit, &refFruit)

}
