package main

import "fmt"

var AccessToken string = "asfkmvv"

func main() {
	var username string = "nikhil"
	fmt.Println(username)
	fmt.Printf("Variables is of type: %T \n",username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variables is of type: %T \n",isLoggedIn)

	var smallValue int16 = 255
	fmt.Println(smallValue)
	fmt.Printf("Variables is of type: %T \n",smallValue)

	var floatValue float32 = .122345566778990
	fmt.Println(floatValue)
	fmt.Printf("Variables is of type: %T \n",floatValue)

	var anotherValue int
	fmt.Println(anotherValue)
	fmt.Printf("Variables is of type: %T \n",anotherValue)

	// implicit type

	var num = 1233
	fmt.Println(num)

	// no var style

	numberOfUsers := 303000
	fmt.Println(numberOfUsers)


	fmt.Println(AccessToken)

}