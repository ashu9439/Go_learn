//go run .\waitGroup.go
// wait till all the parallel tasks are marked as done

package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	var waitSignal sync.WaitGroup

	waitSignal.Add(1)
	go parallelProcess1(&waitSignal, 1)

	waitSignal.Add(1)
	go parallelProcess1(&waitSignal, 2)

	waitSignal.Add(1)
	go parallelProcess1(&waitSignal, 3)

	waitSignal.Wait()

	fmt.Println("all jobs completed")
}

func parallelProcess1(waitSignal *sync.WaitGroup, i int) {
	defer func() {
		waitSignal.Done()
	}()

	fmt.Println("pushed into channel " + strconv.Itoa(i))
}
