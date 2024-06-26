// go run .\multiple_input_tochannel_2.go
package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	channel := make(chan string)

	go parallelProcess(channel, 1)
	go parallelProcess(channel, 2)
	go parallelProcess(channel, 3)
	go parallelProcess(channel, 4)
	go parallelProcess(channel, 5)

	for msg := range channel {
		fmt.Println(msg)
	}

}

func parallelProcess(channel chan string, i int) {
	rand.Seed(time.Now().UnixNano())
	randomSeconds := rand.Intn(3)
	time.Sleep(time.Second * time.Duration(randomSeconds))
	channel <- "pushed into channel " + strconv.Itoa(i) + "  waited for" + strconv.Itoa(randomSeconds)
}
