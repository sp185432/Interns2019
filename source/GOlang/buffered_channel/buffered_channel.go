package main

import (
	"fmt"
	"sync"
	"time"
)

var waitGrp sync.WaitGroup

func send(ele int, buf chan int) {

	value := <-buf
	fmt.Println("Inside the function:", value)
	time.Sleep(1 * time.Second)

	waitGrp.Done()

	buf <- ele //this statement must be after Done(). Otherwise its a Deadlock.
}
func main() {
	buf1 := make(chan int, 3)
	buf1 <- 11
	buf1 <- 22
	buf1 <- 33
	//buf1 <- 66

	//If data more than the size of buffer is added, it causes a deadlock

	fmt.Println("Buffer 1- Length/Number of items in channel BEFORE RECEIVING is: ", len(buf1))

	/*if the no. of iterations and channel length are equal and less
	amount of data is sent then it causes deadlock as the buffer becomes empty*/

	for i := 1; i <= 3; i++ {
		fmt.Println(<-buf1)
	}

	fmt.Println("Buffer 1- Number of items in channel AFTER RECEIVING is: ", len(buf1))

	buf := make(chan int, 1)

	waitGrp.Add(5)

	for i := 1; i <= 5; i++ {
		go send(i, buf)
	}

	buf <- 44

	waitGrp.Wait()

	fmt.Println("Inside the main: ", <-buf)

	fmt.Println("end")
}
