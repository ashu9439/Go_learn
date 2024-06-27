// a pool is a collection of reusable resources. 
// Pools are useful for managing resources that are expensive to allocate and release, 
// such as network connections or large data structures. 
// The sync package provides a sync.Pool type to create and manage pools of objects.

package main

import (
	"fmt"
	// "strconv"
	"sync"
	// "time"
)

var ()

func main() {
	noOfResource := 0
	resourcePool := &sync.Pool{
		New: func() interface{} {
			noOfResource++
			res := make([]byte, 1024)
			return &res
		},
	}



	var waitSignal sync.WaitGroup
	waitSignal.Add(1000)

	for i := 0; i < 1000; i++ {
		go func (){
			res:= resourcePool.Get()
			resourcePool.Put(res)
			waitSignal.Done()
		}()
	}

	waitSignal.Wait()
	fmt.Println("total resource created - ", noOfResource)

}



