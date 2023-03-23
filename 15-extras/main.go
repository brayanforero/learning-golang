// Topic: defer, panic, recover
package main

import "fmt"

func main() {

	// proccesWithDefer()
	fmt.Printf("\n")
	// proccesDeferPills()
	// proccesVariable()
	fmt.Println("Result: ", panicWow(4.5, 2))
	fmt.Println("Result: ", panicWow(4.5, 0))
	fmt.Println("Result: ", panicWow(6.54, 3))
}

// The defer intruction can stop a code sentence and excute it to end
func proccesWithDefer() {

	defer fmt.Println("STEP 1")
	fmt.Println("STEP 2")
	fmt.Println("STEP 3")
}

// The defer intruction add in a execution row and excute this way:
// defer senteceA
// defer senteceB		-> [senteceC,senteceB,senteceA] execute the last added
// defer senteceC

func proccesDeferPills() {

	defer fmt.Println("STEP 1")
	fmt.Println("STEP 2")
	defer fmt.Println("STEP 3")
}

// The defer intruction take the all value and used. does not matter this values gointo change its value
func proccesVariable() {
	x := 5
	defer fmt.Println("STEP 1", x)

	x = 10
	fmt.Println("STEP 2", x)

}

// Panic intruction stop the exuction of code like throw new Error
// This show all lines where happened the error

// Recover can continue with the execution of code and do not stop
// It must use the defer instruction
func panicWow(dividendo float32, divisor float32) float32 {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover from: ", r)
		}
	}()

	if divisor <= 0 {
		panic("Expected a valid value")
	}

	return dividendo / divisor
}
