package main

import (
	"fmt" 
	"strconv"
)

// first letter in Upper case export the function out of the package
func Hello (name string) {
	fmt.Print("Hello, ", name, "!\n")
}


// typed return
func sum (a int, b int) int {
	return a + b
}

func convertAndSum(a int, b string) int {
	// ignore one of the values
	value, _ := strconv.Atoi(b)

	return a + value
}


// more complex function with a return struct defined 
// the return variables no need one syntax like `return result`
func convertStringToInt(value string) (result int, err error) {
	converted, err := strconv.Atoi(value)

	if err != nil {
		return
	}

	result = converted

	return
}

func main () {
	hello("romesdev")

	fmt.Println(sum(21, 21))

	fmt.Print("total: ", convertAndSum(10, "2"))

	result, err := convertStringToInt("221")
	fmt.Println("result: ", result, err)
}