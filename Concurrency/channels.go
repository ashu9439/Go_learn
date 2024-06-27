// go run .\channels.go

// Channels in Go are a way for goroutines to communicate with each other and synchronize their execution. 
// Think of a channel as a pipe through which you can send and receive data.
package main

import (
	"fmt"
	"time"
)

func main() {
	startTime := time.Now()
	defer func() {
		fmt.Println("total time taken = ", time.Since(startTime))
	}()

	/*
		isCompleted := make(chan bool)
		go parralellProcess(isCompleted)

		//* <-isCompleted *: Blocks the main goroutine until it receives a value from the "isCompleted channel"
		fmt.Println("parallel process completed ? - ", <-isCompleted)
	*/

	/*
	   // below code will cause error , since default capacity for a chanell = 0 , i.e. if someting put into the cahnnel by any process then it must be immediately consume by the parent/ main process
	   // but here it is not immediately consuming, instead it is waiting till the executing go till the next line
	   	ch := make(chan string)
	   	ch <- "1st msg"
	   	fmt.Println("data in channel ch - ", <-ch)
	*/

	ch := make(chan string, 2)
	// channels is like a FIFO queue
	ch <- "1st msg"
	ch <- "2nd msg"
	fmt.Println("data in channel ch - ", <-ch)
	fmt.Println("data in channel ch - ", <-ch)
}

func parralellProcess(isCompleted chan bool) {
	time.Sleep(time.Second * 2)
	isCompleted <- true
}
