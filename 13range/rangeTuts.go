package main

import (
	"fmt"
)

func main() {
	// Tuts for Range Function

	mylist := []int{1, 2, 3, 4, 500, 1525, 76}

	mymap := map[string]string{"Hello": "World", "Aadarsh": "Pandey", "Kunal": "Tripathi"}

	// In this Syntax First El is INDEX of the Num & Second INDEX is the Actual value of the Num
	for i, num := range mylist {
		fmt.Println("Index =", i, "Value =", num)
	}

	// In similar fashion we can traverse the Maps Key Values
	for key, val := range mymap {
		fmt.Println("Key =", key, "Value =", val)
	}

	// We can also use it to traverse the String Char by Char
	// Outputs Unicode for that particular characters

	myname := "Aadarsh Pandey"
	for i, char := range myname {
		fmt.Println("Idx =", i, "char =", char, "Value =", string(char))
	}
}
