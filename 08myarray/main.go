package main

import "fmt"

func main() {
	fmt.Println("Welcome to Lessions on Arrays in GoLang")

	// Size  must be declared at the creation time itelf
	var nameList [4]string

	nameList[0] = "Aadarsh"
	nameList[1] = "Kush"
	// fruitList[2] = "Rupesh"
	nameList[3] = "Kunal"

	fmt.Println("Printing the Array=>", nameList)
	fmt.Println("Printing the Array Size=>", len(nameList))
	fmt.Println("Printing the 2nd Index Array Item=>", nameList[2])

	// Inline Array Declaration
	var sirNameList = [4]string{"Pandey", "Jalan", "Tripathi"}
	fmt.Println("SirName List=>", sirNameList)

	// Declaring 2D Arrays inline
	fullNameArray := [3][2]string{{nameList[0], sirNameList[0]}, {nameList[1], sirNameList[1]}, {nameList[3], sirNameList[2]}}

	fmt.Println("Full Name Array =", fullNameArray)

	// =========================== Output ===========================

	// 	Welcome to Lessions on Arrays in GoLang
	// Printing the Array=> [Aadarsh Kush  Kunal] -> { Notice the <space> between
	// 0th idx item and 1st idx item & <space><space> after 1st index item [Kush],
	// suggesting that 2nd index item is misisng }
	// Printing the Array Size=> 4
	// Printing the 2nd Index Array Item=>
	// SirName List=> [Pandey Sharma Tripathi ] -> { Notice the <space> at the end after Tripathi}

	// ============================== USE CASES =====================================

	// - When ArraySize is Known
	// - Constant Time O(1) Access -> Faster Access (Since size & indexes are known)
	// - Memory Optimisation
}
