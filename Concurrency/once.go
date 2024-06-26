// sync.Once type is used to ensure that a piece of code is executed only once, 
// no matter how many times it is called from different goroutines. 
// This is particularly useful for initialization code, 
// such as setting up a singleton or initializing resources.

package main

import (
	"fmt"
	// "strconv"
	"sync"
	// "time"
)

var (
once  sync.Once
)

func main() {
	var waitSignal sync.WaitGroup

	waitSignal.Add(3)

	for i := 0; i < 3; i++ {
		go parallelTask(&waitSignal, i, 1)
	}

	waitSignal.Wait()

}

func parallelTask(waitSignal *sync.WaitGroup, i int, wait int) {
	defer func() {
		waitSignal.Done()
	}()
	fmt.Println("parallel task working", i)

	once.Do(taskThatNeedsToBeDoneOnce)
}

func taskThatNeedsToBeDoneOnce(){
	fmt.Println("performed the task, that needs to be done once")
}