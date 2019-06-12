//"mutex" is used to lock a whole 'block' and "atomic" is used to lock a single 'resource'.
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var ( //declaring program scope variables
	count int32
	waitGrp sync.WaitGroup
)

func inc(name string) {
	defer waitGrp.Done()

	for i:=1; i<= 5; i++ {
		atomic.AddInt32(&count, 1) //increment count by '1'
		//count++
		fmt.Println(name,":",count)
	}
	
}

func main() {
	waitGrp.Add(3)

	go inc("one")
	go inc("two")
	go inc("three")

	waitGrp.Wait()
}