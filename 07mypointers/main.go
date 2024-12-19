package main

import "fmt"

func main() {
	fmt.Println("Hello Welcome to Lession on Pointers in GoLang")

	var thisPtr *int
	fmt.Println("Value of thisPtr is", thisPtr)

	myNumber := 23
	var newPtr = &myNumber
	fmt.Println("Value of myNumber is", *newPtr)
	fmt.Println("Value fof myNumber address is", newPtr)

	// Using Ptrs to Modify The data will change the actual data
	// Example

	*newPtr += 2
	fmt.Println("Modified Value of myNumber =", myNumber)

}
