package main

import "fmt"


type paymenter interface {
	pay (amount float32)
	refund(amount float32,account string)
}

type payment struct{
	gateway paymenter
}

// open close principle
func (p payment) makePayment(amount float32) {
	//razopayPaymentGW := razorpay{}
	//stripePaymentGW := stripe{}
	//razopayPaymentGW.pay(amount)
	p.gateway.pay(amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	// logic to make payment
	fmt.Println("Making payment  using razorpay",amount)
}

// type stripe struct {}

// func (s stripe) pay(amount float32){
// 	fmt.Println("Making payment using stripe",amount)
// }

type fakePayment struct {}

func (f fakePayment) pay (amount float32){
	fmt.Println("making payment using fake payment gateway for testing purpose")
}

type paypal struct {}

func (p paypal) pay(amount float32){
	fmt.Println("Making payment using paypal",amount)
}

func (p paypal) refund(amount float32, account string){
	fmt.Println("Refund to user")
}

func main() {
	//stripePaymentGW := stripe{}
	//razorpayPaymentGW := razorpay{}
	//fakeGW := fakePayment{}
	paypalGW := paypal{}
	newPayment := payment{
		gateway: paypalGW,
	}
	newPayment.makePayment(100)
}