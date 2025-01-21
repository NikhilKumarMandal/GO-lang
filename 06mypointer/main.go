// package main

// import "fmt"

// func main() {
// 	fmt.Println("Welcome to pointer class")

// 	// var ptr *int
// 	// fmt.Println("value of pointer is:", ptr)

// 	myNumber := 25

// 	var ptr = &myNumber
// 	fmt.Println("value of actural pointer is",ptr)
// 	fmt.Println("value of actural pointer is",*ptr)

// 	*ptr = *ptr + 2
// 	fmt.Println("New value is:",myNumber)

// }

package main

import "fmt"

// by value
// func changeNum(num int){
// 	num = 5
// 	fmt.Println("In changeNum",num)
// }

// by refrenece
func changeNum(num *int){
	*num = 5
	fmt.Println("In changeNum",*num)
}

func main(){
	num := 1
	changeNum(&num)
	// changeNum(num)

	// fmt.Println("Memory Address",&num)

	fmt.Println("After changeNum in main func",num)
}