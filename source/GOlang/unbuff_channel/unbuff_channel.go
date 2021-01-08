package main

import (
	"fmt"
	"sync"
	"time"
)

var	waitGrp sync.WaitGroup

//NOTE: If the data is sent into a channel and there is no receiver, the control enters into waiting state till the receiver is activated and this leads to DEADLOCK.

func add(ele int, unbuf chan int) {
//	unbuf <- ele //causes deadlock
	value:= <-unbuf
	fmt.Println("func:",value)
	time.Sleep(1 * time.Second)
	
	waitGrp.Done()

	unbuf <- ele //this statement must be after Done(). Otherwise it is a Deadlock in this code.
}


func main() {
	unbuf := make(chan int)
	//defer close(unbuf) //close the channel until the end of main() execution
	
	waitGrp.Add(3)
  
	// for i:=2; i<=3; i++ {  //deadlock causing code
	// 	go add(i,unbuf)
	// }

	go add(10,unbuf)
	go add(20,unbuf)
	go add(30,unbuf)
	
	unbuf<-40

	waitGrp.Wait()
	
	fmt.Println("main:",<-unbuf)

	fmt.Println("end")
}

