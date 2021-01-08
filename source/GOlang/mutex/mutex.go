//"mutex" is used to lock a whole 'block' and "atomic" is used to lock a single 'resource'.
package main

import (
	"fmt"
	"sync"
)

var ( //declaring program scope variables
	my_mutex sync.Mutex
	waitGrp sync.WaitGroup
)

func inc(name string) {
	defer waitGrp.Done()

	my_mutex.Lock() //opening on next line
	{
		for i:=1; i<= 5; i++ {
			fmt.Println(name,":",i)
		}
	}
	my_mutex.Unlock()
}

func main() {
	waitGrp.Add(2)

	go inc("one")
	go inc("two")

	waitGrp.Wait()
}