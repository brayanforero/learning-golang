// Topic: Loops
package main

import "fmt"

func main() {

	// 1. FOR
	fmt.Printf("FOR\n")
	for i := 0; i < 10; i++ {
		fmt.Println("Value: ", i)
	}

	// 2. WHILE
	fmt.Printf("\nWHILE\n")
	i := 0
	for i < 6 {
		fmt.Println("Count: ", i)
		i++
	}

	// 3.DO WHILE
	fmt.Printf("\nDO WHILE\n")
	value := 0
	for next := true; next; next = value < 5 {
		fmt.Println("DO while: ", value)
		value++
	}

}
