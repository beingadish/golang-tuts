package main

import (
	"fmt"
	"sync"
	"time"
)

func processNum(numChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second)
	// Stream of Nums
	fmt.Println("Receiving Nums from Channel <-", <-numChan)
}

func sumChan(result chan int, a, b int) {
	sum := a + b
	result <- sum
}

func main() {

	numChan := make(chan int)
	var wg sync.WaitGroup

	// Below Code Processes Smoothely without Any Blocking
	go sumChan(numChan, 5, 5)

	res := <-numChan

	fmt.Println("Res =", res)

	// ======================= CODE FOR PROCESSING CONTIGUOUS NUMBER OF STREAMS USING WAITGROUPS & FOR LOOP =========================================

	// for i := range 100 {
	// 	wg.Add(1)
	// 	go processNum(numChan, &wg) // Non Blocking (due to go, runs on another thread)
	// 	numChan <- i
	// }

	// Closing the channel before ending the Program
	close(numChan)

	wg.Wait()
	// Making a Message Channel using make() method, we use "chan" keyword for that
	// -> here we have created a channel for passing only STRING messages
	// messagechan := make(chan string)

	// these are blocking in nature (a kind of PIPE) // Cant send data unless another side is ready for receiving data
	// messagechan <- "ping" // (SENDING) => Here "Ping" message is going inside the messageChannel
	// -> meaning we are feeding the message to the channel (we want it to send to other side)

	// msgReceived := <-messagechan // RECEIVING DATA

	// Creates for -> {DEADLOCK!!}
	//      fatal error: all goroutines are asleep - deadlock!

	// goroutine 1 [chan send]:
	// main.main()
	//         /mnt/d/CODE/Go/24channels/channels.go:11 +0x36
	// exit status 2
	// fmt.Println(msgReceived)

}
