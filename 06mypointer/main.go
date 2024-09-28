package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointer class")

	// var ptr *int
	// fmt.Println("value of pointer is:", ptr)

	myNumber := 25

	var ptr = &myNumber
	fmt.Println("value of actural pointer is",ptr)
	fmt.Println("value of actural pointer is",*ptr)

	*ptr = *ptr + 2
	fmt.Println("New value is:",myNumber)

}