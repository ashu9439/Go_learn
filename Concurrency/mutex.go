//go run .\waitGroup.go
// wait till other process release the lock and proceed only when aquair the lock

package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var (
	lock  sync.Mutex
	rlock sync.RWMutex
	count int
)

func main() {
	count = 0
	var waitSignal sync.WaitGroup

	waitSignal.Add(6)
	go parallelProcessRead(&waitSignal, 1, 1)
	go parallelProcess2(&waitSignal, 1, 3)
	go parallelProcessRead(&waitSignal, 2, 1)
	go parallelProcessRead(&waitSignal, 3, 1)
	go parallelProcess2(&waitSignal, 2, 2)
	go parallelProcess2(&waitSignal, 3, 2)

	waitSignal.Wait()

	fmt.Println()
	fmt.Println("all jobs completed, Total count = ", count)
}

func parallelProcess2(waitSignal *sync.WaitGroup, i int, wait int) {
	lock.Lock()
	fmt.Println("locked job " + strconv.Itoa(i))
	defer func() {
		waitSignal.Done()
		lock.Unlock()
		fmt.Println("Unlocked job " + strconv.Itoa(i))
	}()
	count = increase(count, wait)
}
func increase(a int, wait int) int {
	time.Sleep(time.Second * time.Duration(wait))
	a++
	return a
}

func parallelProcessRead(waitSignal *sync.WaitGroup, i int, wait int) {
	lock.Lock()
	fmt.Println("read locked job " + strconv.Itoa(i))
	defer func() {
		waitSignal.Done()
		lock.Unlock()
		fmt.Println("read Unlocked job " + strconv.Itoa(i))
	}()
	fmt.Println("value read = " + strconv.Itoa(count))
}


