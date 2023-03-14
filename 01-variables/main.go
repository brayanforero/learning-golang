package main

import "fmt"

func main() {

	// TOPIC: VARIABLES
	// How to declare it

	// 1. var myVar [string, int, bool ,....] = value
	var myName string = "Brayan"
	fmt.Println(myName)

	// 2. var myVar string -> It is only declaration
	// myVar = value -> Then assigment it a value

	var myFather string
	myFather = "Bricher"
	fmt.Println(myFather)

	// 3. var myVar = value -> This case the laguaje by itself assgiment it data type

	var myMom = "Luz Maira"
	fmt.Println(myMom)

	// 4. myVar := value -> This case is common on profesional enviroments, it is about declaration and set value in the variables
	myBrother := "Roswell"
	fmt.Println(myBrother)

	fmt.Println("This is my family:")
	fmt.Printf("*Father: %s\n", myFather)
	fmt.Printf("*Mom: %s\n", myMom)
	fmt.Printf("*Brother: %s\n\n", myBrother)
	fmt.Printf("See you next, Bye bye :)")

}
