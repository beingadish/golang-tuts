package main

import "fmt"

func myTask(done chan bool) {
	// Using Anonymous Function to send True via channel that the task is done at the end
	// (end because we're using <defer> keyword)
	defer func() { done <- true }()
	for i := range 999999999999 {
		fmt.Println("Processing Task", i)
	}
}

// Error because Two main() functions are defined, but its Okay if we
// run this file independently because then it wont create Ambiguity
func main() {
	// Another Channel Tuts for replacements of WaitGroup

	done := make(chan bool)

	// This allows go to process Task in another Thread
	go myTask(done)

	// Here we are trying to receive the BOOLEAN which will be sent
	// through channel after the myTask() gets processed
	<-done // Block until data does not comes
}
