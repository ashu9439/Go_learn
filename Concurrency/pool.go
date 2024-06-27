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



