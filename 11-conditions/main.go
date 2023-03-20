// Topics: Conditions
package main

import "fmt"

func main() {

	// 1. Firts way
	message := "Cannot get in"
	userAge := 26

	if userAge < 18 {
		message = "You are child yet"
	} else if userAge > 45 {
		message = "Your are old man, cannot in"
	} else {
		message = "Welcome to my app, Sir"
	}

	fmt.Println(message)

	// 2. Second way sort

	if things := "🍞"; things == "🍎" {
		fmt.Println("The things is a fruit")
	} else if things == "🍞" {
		fmt.Println("The things is a cereal")
	} else {
		fmt.Println("Sorry, I do not know what it is")
	}
}
