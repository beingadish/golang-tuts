package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter your X Std. CGPA: ")
	reader := bufio.NewReader(os.Stdin)

	cgpa, _ := reader.ReadString('\n')

	percentage, err := strconv.ParseFloat(strings.TrimSpace(cgpa), 64)

	if err != nil {
		fmt.Println("Type Conversion Went Wrong !!")
	} else {
		fmt.Println("Your X Std, Percentage =", percentage*9.5)
	}

}
