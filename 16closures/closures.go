package main

import "fmt"

// Closure Function

func counter() func() int {
	var count int = 0

	return func() int {
		count++
		return count
	}
}

func main() {
	// Golang Tuts on Closures

	myFun := counter()

	fmt.Println(myFun())
	fmt.Println(myFun())
	fmt.Println(myFun())
	fmt.Println(myFun())
	fmt.Println(myFun())
}
