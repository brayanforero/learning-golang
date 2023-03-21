// Topis: fuctions
package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

// Normal function
func sayHello() {
	fmt.Println("Hello world from Golang")
}

// Functions with paramaters
// fun name (parameter datatypoe) {body}
// If all paramatares are equlas datatype, it is can put the datatype to end
// fun name (num1 ,num2 ,num3 int) {body}
func sayHelloWithYourName(name string) {
	fmt.Printf("Hellow world %s from Golang", strings.ToUpper(name))
}

// Functions with returned multiple values
// fun name (parameters? datatype) datatype {value}
func sum(a, b int) int {

	return a + b
}

// Functions with returned valeu
// fun name (parameters? datatype) (datatypeOne, dataTypeTwo) {value}
func transformText(firstname string) (string, string) {

	textA := strings.ToUpper(firstname)
	textB := strings.ToLower(firstname)

	return textA, textB
}

func goinToReadFile(source string) {

	file, err := ioutil.ReadFile(source)

	if err != nil {
		fmt.Println("Has been an error reading your file: ", err)
		return
	}

	fmt.Println(file)
}

func generateCustomError(number int) (result int, err error) {

	if number <= 0 || number > 10 {
		err = errors.New("Expected a value beetwen 1 and 10")
		return
	}

	result = number * 100
	return
}

func filterFromArray(array []int, callback func(item int) bool) []int {
	result := []int{}

	for _, v := range array {

		if callback(v) {

			result = append(result, v)
		}

	}

	return result
}
// Like rest operator in JavaScript
func sumAllNumber(nums ...int) int {
	result := 0

	for _, v := range nums {
		result += v
	}

	return result
}

func main() {

	// sayHello()

	// sayHelloWithYourName("brayan")

	// result := sum(4, 5)

	// fmt.Println("Result:", result)

	// a, b := transformText("brayan")

	// fmt.Println("Result: ", a, ", ", b)

	// goinToReadFile("hello.txt")

	// result, err := generateCustomError(5)

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println(result)

	// nums := []int{1, 2, 3, 4, 5, 34, 50, 123}
	// fmt.Println("Original Array ", nums)
	// newArray := filterFromArray(nums, func(item int) bool {
	// 	return item%2 != 0
	// })
	// fmt.Println(newArray)

	fmt.Println(sumAllNumber(1, 2, 3, 4))

}
