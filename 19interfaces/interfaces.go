package main

import "fmt"

// Defining an Interface

type payementGateway interface {
	pay(amount float64)
}

// Creating Gateways by following the same method_signatures of interface
// ==============================================================
type fakePay struct{}

func (f fakePay) pay(amount float64) {
	fmt.Println("Payment done using FakePay of Amount =", amount)
}

// func (f fakePay) payAnother(amount float64) {
// 	fmt.Println("Payment done using FakePay of Amount =", amount)
// }

// ==============================================================
type PhonePay struct{}

func (p PhonePay) pay(amount float64) {
	fmt.Println("Payment done using PhonePay of Amount =", amount)
}

// ==============================================================
type GooglePay struct{}

func (g GooglePay) pay(amount float64) {
	fmt.Println("Payment done using GooglePay of Amount =", amount)
}

// ==============================================================
type ApplePay struct{}

func (a ApplePay) pay(amount float64) {
	fmt.Println("Payment done using ApplePay of Amount =", amount)
}

// ================== Creating a Payment Client Struct To Make Payments ==========================
type paymentClient struct {
	payementGateway
}

func (p paymentClient) makePayment(amount float64) {
	p.payementGateway.pay(amount)
}

// Implementing them in Main()

func main() {
	// Golang Interfaces Tuts

	// =-=-=-=-=-=-=-=-=- Creating Gway Instances -=-=-=-=-=-=-=-=-=-=
	applePayGw := ApplePay{}

	gpayGw := GooglePay{}

	phonepayGw := PhonePay{}

	fakePayGw := fakePay{}

	// Payment Using All of them One by One

	// =-=-=-=-=-=-=-=-=- Creating Diff Clients for Diff GWay Instances -=-=-=-=-=-=-=-=-=-=
	applePayClient := paymentClient{applePayGw}
	gPayClient := paymentClient{gpayGw}
	phonePayClient := paymentClient{phonepayGw}
	fakePayClient := paymentClient{fakePayGw}

	// Initiating Payments

	applePayClient.makePayment(100)
	gPayClient.makePayment(1000)
	phonePayClient.makePayment(999)
	fakePayClient.makePayment(852478)

}
