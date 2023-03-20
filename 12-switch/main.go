// Topic: Switch
package main

import "fmt"

func main() {

	const COLOR = "yellow"

	// 1. first way

	switch COLOR {
	case "yellow":
		fmt.Println("Descrease your velocity")
	case "red":
		fmt.Println("STOP")
	default:
		fmt.Println("GO")
	}

	// 1. second way

	const age = 12

	switch {
	case age < 30:
		fmt.Println("YOUNG")
	case age >= 30 && age <= 60:
		fmt.Println("ADULT")
	default:
		fmt.Println("OLD MAN")
	}
}
