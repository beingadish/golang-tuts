package main

import (
	"fmt"
	"time"
)

func doTask(id int) {
	fmt.Println("Performing Task =", id)
}

func main() {
	// Golang Tuts on Goroutines -> Lightweight Threads (used for multi-threading)

	for i := range 101 {
		// => Acts as a Non-Blocking Task
		go doTask(i) // -> Performing Go Routines very fast 
		// -> goes to scheduler & then the scheduler processes the Threads parallelly
	}

	time.Sleep(time.Second)
}
