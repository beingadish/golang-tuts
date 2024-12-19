package main

import (
	"fmt"
	"time"
)

// Defining a Struct

type OrderStatus struct {
	name string
}

type OrderEntity struct {
	id          int
	amount      float64
	createdAt   time.Time // -> Nanosecond Precision
	desc        string
	OrderStatus // -> Struct Embedding
}

// HACK for Constructor of Structs
func newOrderEntity(id int, amount float64, desc string) *OrderEntity {

	/// Preprocessing Goes here ...
	orderEntity := OrderEntity{
		id:        id,
		amount:    amount,
		desc:      desc,
		createdAt: time.Now(),
		OrderStatus: OrderStatus{
			name: "New",
		},
	}

	return &orderEntity
}

// Using Receiver Type to Change the OrderStatus (can be used for creating Setters & Getters)
func (o *OrderEntity) changeOrderStatus(orderStatus OrderStatus) {
	o.OrderStatus = orderStatus
}

func main() {
	// Golang Tuts on Structs

	// Creating a memory instance of the OrderEntity Struct
	order1 := OrderEntity{
		id:        21,
		amount:    852.0,
		createdAt: time.Now(),
		desc:      "Hello Some Description",
		OrderStatus: OrderStatus{
			name: "New",
		},
	}

	order2 := newOrderEntity(14, 5883.22, "Order Entity 2")

	fmt.Println("Printing Order Entity 2 =", order2)

	currStatus := OrderStatus{
		name: "In Transit",
	}

	fmt.Println("Order1 Entity Before InTransit =", order1)
	order1.changeOrderStatus(currStatus)
	fmt.Println("Order1 Entity After InTransit =", order1)

	// ============ Creating a Struct that is to be used only one ====================

	language := struct {
		name   string
		isGood bool
	}{
		"golang",
		true,
	}

	fmt.Println("Temp Language Struct =", language)
}
