package main

import "fmt"

// Passing values as a reference (so actual value changes)
func incrementOne(a *int) {
	// * before Pointer argument to Dereference the Pointers
	*a++
}

// A copy of value is being Passed
func withoutReferenceIncrementOne(b int) {
	b++
}

func main() {

	var a int = 0
	var b int = 0
	fmt.Println("Memory Address of A =", &a, "& Address of B =", &b)
	fmt.Println("Before Increment =", a, b)
	incrementOne(&a)                       // -> Leads to Increment ONEs
	withoutReferenceIncrementOne(b)        // -> Leads to Nothing
	fmt.Println("After Increment =", a, b) // -> Prints 1 0

}
