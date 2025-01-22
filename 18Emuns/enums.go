package main

import "fmt"

// Enumerated types

type OrderStatus int 

const (
	Received OrderStatus = iota
	Confirmed 
	Prepared
	Delivered
)

func chnageOrderStatus(status OrderStatus){
	fmt.Println("Changing order status to",status)
}

func  main()  {
	chnageOrderStatus(Delivered)
}