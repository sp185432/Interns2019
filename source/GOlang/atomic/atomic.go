package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var count int32
var waitGrp sync.WaitGroup

func increment(name string) {
	defer waitGrp.Done()

	for i := 1; i <= 5; i++ {
		atomic.AddInt32(&count, 4) //increment count by '4'

		fmt.Println(name, ":", count)
	}

}

func main() {
	waitGrp.Add(3)

	go increment("one")
	go increment("two")
	go increment("three")

	waitGrp.Wait()
}
