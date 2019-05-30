package main

import (
	"fmt"
	"sync"
	"time"
)

var waitGrp sync.WaitGroup

/* In a buffered channel, when a sender sends data there should be a receiver to receive
otherwise the control enters the waiting state till the receiver becomes available and it causes deadlock*/

func send(ele string, unbuf chan string) {
	//	unbuf <- ele //causes deadlock
	data := <-unbuf
	fmt.Println("Inside the go function : ", data)
	time.Sleep(1 * time.Second)

	waitGrp.Done()

	unbuf <- ele //after its once this statement should be executed, otherwise it causes deadlock
}

func main() {
	unbuf := make(chan string)
	//defer close(unbuf) //close the channel until the end of main() execution

	waitGrp.Add(5)

	go send("This", unbuf)
	go send("program", unbuf)
	go send("tells the way how ", unbuf)
	go send("buffered channels work in", unbuf)
	go send("GO language", unbuf)
	//go send("haii", unbuf)
	//go send("hello", unbuf)

	unbuf <- "Communication between GoRoutines: CHANNELS"

	waitGrp.Wait()

	fmt.Println("Inside the main: ", <-unbuf)

	fmt.Println("End of the program")
}
