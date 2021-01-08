package main

import (
	"fmt"
	"sync"
	"time"
)

var	waitGrp sync.WaitGroup

//NOTE: If the data is sent unbuffered channel and there is no receiver, the control enters into waiting state till the receiver is activated and this leads to DEADLOCK.

func add(ele int, buf chan int) {
	
	value:= <-buf
	fmt.Println("func:",value)
	time.Sleep(1 * time.Second)

	waitGrp.Done()

	buf <- ele //this statement must be after Done(). Otherwise its a Deadlock.
}


func main() {
	buf1 := make(chan int, 3)
	buf1 <- 1
	buf1 <- 2
	buf1 <- 3

//If data more than the size of buffer is added, it causes a deadlock

fmt.Println("Buffer 1- Length/Number of items in channel BEFORE RECEIVING is: ",len(buf1))

for i:=1; i<=3; i++ {
	fmt.Println(<-buf1) 
}	

fmt.Println("Buffer 1- Number of items in channel AFTER RECEIVING is: ",len(buf1))
	

buf2 := make(chan string, 10)
	buf2 <- "hello"
	buf2 <- "welcome"
	buf2 <- "hi"
	buf2 <- "fine"
fmt.Println("Buffer 2- Length/Number of items in channel BEFORE RECEIVING is: ",len(buf2))

for i:=1; i<=4; i++ {
	fmt.Println(<-buf2) 
}	

fmt.Println("Buffer 2- Number of items in channel AFTER RECEIVING is: ",len(buf2))
	

	buf := make(chan int, 1)
	waitGrp.Add(3)

	for i:=1; i<=3; i++ {
		go add(i,buf)
	}
	
	// go add(10,buf) //if we remove 'go' then multithreading is stopped and it results in deadlock as the buffer has no elements as of now and in the "add()" function, the first statement is receive.
	// go add(20,buf)
	// go add(30,buf)

	 buf<-40
	
	waitGrp.Wait()
	
	fmt.Println("main:",<-buf)

	fmt.Println("end")
}

