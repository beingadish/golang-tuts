package main

import "fmt"

// Enums -> Enumerations => Enumerated Types

// OrderStatus Enum
type OrderStatus int

const (
	NEW OrderStatus = iota
	IN_TRANSIT
	DISPATCHED
	DELIVERED
)

// MessageStatus Enum
type MessageStatus string

const (
	SENT              MessageStatus = "Sent"
	MESSAGE_DELIVERED               = "Delivered"
	READ                            = "Read"
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("Order Status Changed to", status)
}

func changeMessageStatus(messageStatus MessageStatus) {
	fmt.Println("Message Status Changed to", messageStatus)
}

func main() {
	// Golang Enums Tuts

	changeOrderStatus(DELIVERED)
	changeMessageStatus(MESSAGE_DELIVERED)

}
