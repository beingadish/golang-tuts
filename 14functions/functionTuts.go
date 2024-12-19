package main

import "fmt"

// General Function Declaration
func myFunc(a int, b int) int {
	return a + b
}

// Declaring Similar Arg Types at One Time
func myFunc2(a, b int) int {
	return a + b
}

// Void Returning Fn
func myFunc3(a, b int) {
}

// Returning Multiple Type Values via Func
func multiReturn(allowed bool) (string, int, bool) {
	return "adb", 10, allowed
}

// Function accepting another function as its argument
func myDiffFunc(fn func(a bool) (string, int, bool)) (string, int, bool) {
	return fn(true)
}

func main() {
	// Golang Tuts for Functions

	// Trying with True
	var1, var2, var3 := multiReturn(true)

	if var3 {
		fmt.Println("Allowed with var1 =", var1, "var2 =", var2)
	} else {
		fmt.Println("Not Allowed")
	}

	// Not using a particular value then represent it using (_) <underscore>
	vari1, vari2, _ := multiReturn(false)

	if vari2 >= 10 {
		fmt.Println("Allowed with vari1 =", vari1, "vari2 =", vari2)
	} else {
		fmt.Println("Not Allowed")
	}

	// Use case for Function as an Argument inside another function

	myDiffFunc(multiReturn)

}
