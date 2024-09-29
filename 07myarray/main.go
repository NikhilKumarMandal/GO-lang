package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in go")

	var fruitList [4]string

	fruitList[1] = "apple"
	fruitList[2] = "banana"
	fruitList[3] = "orange"

	fmt.Println("Fruit List is ", fruitList)
	fmt.Println("Fruit List is ", len(fruitList))


	var vegList = [4]string{"potato","benas","mashroom"}
	fmt.Println("Vegy List is",vegList)
}