package main

func main() {

	// Implementing While loop using [For] Construct
	// i := 1
	// for i <= 5 {
	// 	fmt.Println("Hello World from Golang", i, "Times")
	// 	i++
	// }

	// Infinite Loop
	// for {
	//     println(i)
	// }

	// Classic For Loop in Golang using Breaks & Continue
	// for i := 0; i <= 3; i++ {
	// 	if i == 2 {
	// 		continue
	// 	}
	// 	fmt.Println("Printing i", i)
	// }

	// Ranges in Golang 1.22

	for i := range 3 {
		println("Printing i =", i)
	}

}
