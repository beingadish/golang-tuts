package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("TimeStudy on GoLang")

	// Accessing Time
	currTime := time.Now()

	//  Formatting Time
	fmt.Println("Current Time = ", currTime.Format("Mon 02-01-2006 15:04:05"))

	// Craeting custom Date/Time
	myBirthDate := time.Date(2001, time.April, 13, 00, 00, 00, 00, time.Local).Format("Mon 02/01/2006")
	fmt.Println("My B'Day is on ", myBirthDate)
}
