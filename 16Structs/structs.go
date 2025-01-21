package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")

	// no inheritance in golang; No super or parent
	
	nikhil := User{"Nikhil","nikhil@go.dev",true,16}
	fmt.Println(nikhil)
	fmt.Printf("hites details are:%+v\n",nikhil)
}

type User struct {
	Name string
	Email string
	Status bool
	Age int
}

