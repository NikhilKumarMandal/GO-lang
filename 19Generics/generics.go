package main

import "fmt"

func printSlice[T comparable, V string](items []T,name V) {
	for _, item := range items {
		fmt.Println(item,name)
	}
}

// func printStringSlice(items []string) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// type Stack[T any] struct {
// 	elements []T
// }

func main() {

	// myStack := Stack[int] {
	// 	elements: []int{1,2,3},
	// }

	// fmt.Println(myStack)

	// nums := []int{1,2,3}
	name := []string{"gloang","typescript"} 
	//val := []bool{true,false,true}
	// printStringSlice(name)
	printSlice(name,"nikhil")
}