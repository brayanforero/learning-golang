// TOPIC: maps
package main

import "fmt"

func main() {

	// Maps are data structes that be formed by couples key:value like literal object in javascript

	days := map[string]uint8{
		"MONDAY":   1,
		"TUESDAY":  2,
		"WENSDAY":  3,
		"THURSDAY": 4,
		"FRIDAY":   5,
		"SATURDAY": 6,
		"SUNDAY":   7,
	}

	day, ok := days["MONDAY"]

	if ok {
		fmt.Println("Days of week", day)
		return
	}

	fmt.Printf("The current day %d  is wrong", day)

}
