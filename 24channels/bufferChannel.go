package main

import "fmt"

// CHANNELS act like Queue Pipe (where you can feed data very fast & process asynchronously in diff GoRoutine)

func main() {

	//This will be Non-Blocking in Nature
	// 100 is the Buffer Size of the created channel
	thisBufferChan := make(chan int, 100)

	// If we try to add more values in the channel than that of its size then it leads to a deadlock
	// Program tries to send more values than that of size of Buffer & there is not any receiver present

	thisBufferChan <- 10
	thisBufferChan <- 10
	thisBufferChan <- 10
	thisBufferChan <- 152

	// If we dont close this, then the for loop tries to get the access of integers from the channel
	// but there are no integer available hence a DEADLOCK
	// Deadlock: Happens because the program waits indefinitely for a signal that never comes (i.e., a value or a channel close).
	close(thisBufferChan)

	// This is non-blocking, that's why it waits untill the channel is
	for i := range thisBufferChan {
		fmt.Println("Received", i)
	}

}
