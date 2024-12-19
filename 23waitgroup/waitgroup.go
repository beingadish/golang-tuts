package main

import (
	"fmt"
	"sync"
)

// Accepting WaitGroup Pointer
func doTask(id int, w *sync.WaitGroup) {
	// defer --> keyword used to perform the operation at the very end of the complete execution of function
	// Done() --> clears (-1) from the waitgroup after the task is being executed completely
	defer w.Done()
	fmt.Println("Performing Task =", id)
}

func main() {
	// Golang Tuts on WaitGroups -> Used to Sync the GoRoutines (because in Real-Life Scenario we dont use time.Sleep)

	// Creating an WaitGroup
	var waitgroup sync.WaitGroup

	for i := range 101 {
		// Before running the doTask() method we add 1 to WaitGroup for that doTask()
		waitgroup.Add(1)
		// Passing the WaitGroup Reference Address so it can modify the same WaiGroup()
		go doTask(i, &waitgroup)
	}

	// Wait (Block the Main GoRoutine) --> Till this waitGroup is now being finished (value == 0)
	waitgroup.Wait()
}
