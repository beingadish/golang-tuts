package main

import "fmt"

// Defining a Variadic Function using ellipsis operator

// Nums acts as a Slice
func myIntSum(nums ...int) int {
	sum := 0

	// i = idx && val = actual_value
	for _, val := range nums {
		sum += val
	}

	return sum
}

func main() {
	// Golang Tuts on Variadic Functions (A function which accepts N number of Arguments)

	fmt.Println(myIntSum(5, 47, 9, 5, 2, 52, 1))
	mynums := []int{2, 3, 4, 6, 8, 9, 67}

	// Passing Slices / Arrays using ellipsis operator
	// -> Expanding the slice into individual arguments
	fmt.Println(myIntSum(mynums...))
}
