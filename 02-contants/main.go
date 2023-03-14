package main

import "fmt"

func main() {
	fmt.Println("WELCOME TO THE CONSTANS SECTION")

	// How to use
	// const myConst [string, int, bool, ...] = value
	// const myConst = value
	// cannot use the reversed word ":=" to assign anyvalue

	const PI = 3.1416
	fmt.Printf("This is value of PI: %f", PI)
}
