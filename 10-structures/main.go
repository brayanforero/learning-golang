// Topic: Structures
package main

import "fmt"

func main() {

	// They are data structure can allow to create a new custom data types as POO

	type Person struct {
		name string
		age  uint8
		rol  string
	}

	type Pest struct {
		name       string
		typeOfPest string
		owner      Person
	}

	brayan := Person{
		name: "Brayan",
		age:  24,
		rol:  "Frontend developer",
	}

	flipper := Pest{"Flipper", "dog", brayan}

	fmt.Println("Person: ", brayan.name)
	fmt.Println("\n\nPerson: ", flipper)
}
