package main

import (
	"fmt"
	"maps"
)

func main() {
	// Welcome to Maps in Golang Tuts

	// Creating a MAP using make() keyword [<KEY>]<VALUE>
	mymap := make(map[string]int)

	// Setting elements
	mymap["Hello"] = 25
	mymap["Aadarsh"] = 10
	mymap["Ramesh"]++ // By Default all values are ZERO VALUES
	// --> (<int> 0,<string> "",<bool> false), incrementing it makes it equal to 1 in case of int

	fmt.Println("Printing MyMap As It Is =", mymap)
	fmt.Println("MyMap[Aadarsh] =", mymap["Aadarsh"])
	fmt.Println("MyMap[Ramesh] =", mymap["Ramesh"])

	// Deleting a Key

	delete(mymap, "Aadarsh")
	fmt.Println("Accessing after deleting MyMap[Aadarsh] =", mymap["Aadarsh"])

	// Another way of Making Map in Golang
	mySecondMap := map[string]string{"Hello": "World", "Aadarsh": "Pandey"}

	fmt.Println("This is SecondMyMap[Hello] =", mySecondMap["Hello"], "from Aadarsh", mySecondMap["Aadarsh"])

	// Using a Comma,OK Syntax to check if Key Exists (Good Practice)

	val, OK := mySecondMap["Aadarsh"]

	if OK {
		fmt.Println("Key Exists with Val =", val)
	} else {
		fmt.Println("Key Does Not Exists")
	}

	// maps Package can be used for various methods to be used in Maps

	// fmt.Println(maps.Equal(mymap,mySecondMap)); --> Compile Time Error (Because <Key,Value> Pair does not matches)

	myAnotherMap := map[string]string{"Hello": "World"}

	fmt.Println("Comparing Two Maps using maps.Equal =", maps.Equal(myAnotherMap, mySecondMap))

}
