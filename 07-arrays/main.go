// Topic: Arrays
package main

import "fmt"

func main() {

	// Only can save fixed positions, cannot add new positions

	// 1. Firts way
	var pets [3]string
	pets[0] = "Dog"
	pets[1] = "Cat"
	pets[2] = "Rabbit"

	fmt.Println(pets)

	// 2. Second way
	family := [4]string{"Bricher", "Luz", "Brayan", "Roswell"}

	fmt.Println(family)
}
