package main

import (
	"fmt"
	"sync"
)

var mymutex sync.Mutex
var waitGrp sync.WaitGroup

func inc(name string) {
	defer waitGrp.Done()

	mymutex.Lock() //opening on next line
	{
		for i := 1; i <= 5; i++ {
			fmt.Println(name, ":", i)
		}
	}
	mymutex.Unlock()
}

func main() {
	waitGrp.Add(2)

	go inc("one")
	go inc("two")

	waitGrp.Wait()
}
