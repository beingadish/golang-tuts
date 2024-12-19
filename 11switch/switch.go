package main

import "fmt"

func main() {
	// Normal Switch Case

	// MultiSwitch Case

	// Type Switch Case

	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case string:
			fmt.Println(t, "is a String")
		case int:
			fmt.Println(t, "is an Integer")
		case bool:
			fmt.Println(t, "is a Boolean")
		case float32:
			fmt.Println(t, "is a Float 32")
		case float64:
			fmt.Println(t, "is a Float 64")
		default:
			fmt.Println(t, "is something else")
		}
	}
	var myFloat float32 = 3.14
	whoAmI("Aadarsh")
	whoAmI(10)
	whoAmI(myFloat)
	whoAmI(true)
	whoAmI(234578909874865478632.4785462184785)
}
