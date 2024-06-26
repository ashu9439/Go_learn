// go run .\multiple_input_tochannel.go
package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	startTime := time.Now()
	defer func() {
		fmt.Println("total time taken = ", time.Since(startTime))
	}()

	channel := make(chan string)

	go parallelProcess(channel)

	for msg := range channel {
		fmt.Println(msg)
	}

}

func parallelProcess(channel chan string) {

	for i := 0; i < 5; i++ {
		time.Sleep(time.Second * 1)
		channel <- "pushed into channel " + strconv.Itoa(i)
	}
	close(channel)
}
