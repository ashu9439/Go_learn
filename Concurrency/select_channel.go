// select any one channel
// go run .\multiple_input_tochannel_2.go
package main

import (
	"fmt"
	// "strconv"
	// "time"
	// "math/rand"
)

func main() {
	channel1, channel2 := make(chan string), make(chan string)

	go parallelProcess(channel1,"1st channel returned")
	go parallelProcess(channel2,"2nd channel returned")


	select{
	case msg := <- channel1:
			fmt.Println(msg)
		case msg := <- channel2:
			fmt.Println(msg)
		default:
			fmt.Println("nothing returned")
	}
}

func parallelProcess(channel chan string, i string) {
	channel <- i

	// rand.Seed(time.Now().UnixNano())
    // randomSeconds := rand.Intn(3)
	// time.Sleep(time.Second * time.Duration(randomSeconds))
	// channel <- i + "  waited for" +  strconv.Itoa(randomSeconds)
}
