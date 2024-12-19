package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println("Wecome to Lessions on Slices in GoLang")

	// Not Array but a Slice (Dynamic Array) -> most used construct in golang
	// Provides useful methods to be used onto them

	// Uninitialised Slice is = nil, Size (len(<slice_name>)) = 0
	var fruitList = []string{"apple", "peach", "guava"}

	fmt.Println("FruitList =>", fruitList)

	fruitList = append(fruitList, "pineapple", "mango")

	fmt.Println("This is updated Fruit List=>", fruitList)

	// Slicing the Slice in GoLang
	fruitList = fruitList[1:]

	fmt.Println("This is updated Fruit List=>", fruitList)

	// Allocating SLices using make() Keyword

	// gives slice of int of size 4
	// have 3 params (slice type, initial_length, initial_capacity) -> initial_capacity is by default = initial_length provided
	highScores := make([]int, 4)

	highScores[0] = 524
	highScores[1] = 755
	highScores[2] = 888
	highScores[3] = 999
	// highScores[4] = 1000 // Will Panic [panic: runtime error: index out of range [4] with length 4]

	// But if we use append then its going to reallocate memory &
	// all the values that are being passed
	// using the append() method will get accomodated

	highScores = append(highScores, 7851, 752, 95, 752)
	fmt.Println("HighScore Values=>", highScores)
	slices.Sort(highScores)

	/*
	*Slices Package have good amount of methods that can be used for ex :
	* - slices.Equal(slice1,slice2) ==>  To compare Slices
	* - copy() method for copying the elements from one slice to another slice
	 */

	fmt.Println("HighScore Sorted Values=>", highScores)
}
