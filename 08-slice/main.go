// This topic is about slices
package main

import "fmt"

func main() {

	// What are they Slice
	// They are dinamic arrays can contain multiples items

	//foods := [6]string{"🍇", "🍎", "🍊", "🥑", "🥕", "🥒"}
	//fmt.Println("FOODS: \n", foods)
	// Taking reference of food array
	// array [start:end] -> "end" index is excluded,
	// [,end] -> By default if we not set start index this will be 0
	// [0,] -> By default if we not set end index this will be the last index
	//fruits := foods[0:3] // -> this value is taked by reference, as in the array of Javascript
	// vegetables := foods[3:]
	// fruits[0] = "🍌"

	// fmt.Println("FOODS AFTER CHANGE FIRST ELEMENT: \n", foods)
	// fmt.Println("\nFRUITS: \n", fruits)

	// Now we are going to added a new fruits

	// fruits = append(fruits, "🍓")

	// fmt.Println("\nFOODS AFTER ADDED A 🍓: \n", foods)
	// fmt.Println("FRUITS: ", fruits)

	// fruits = append(fruits, "🥝")
	// fmt.Println("\nFOODS AFTER ADDED A 🥝: \n", foods)
	// fmt.Println("\nFRUITS: \n", fruits)

	// fmt.Println("\nVETEGABLES: ", vegetables)

	bluethings := make([]string, 0)

	bluethings = append(bluethings, "🐦")
	bluethings = append(bluethings, "🐬")

	fmt.Printf("Detalles del Slice de cosas azueles: \n")

	fmt.Println("\n\nItems: ", bluethings)
	fmt.Println("\n\nLogitud: ", len(bluethings))
	fmt.Println("\n\nCapacidad: ", cap(bluethings))

	// When slices are full, this duplicate its length and capacity
}
