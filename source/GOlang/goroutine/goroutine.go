//concurrency or multithreading (achieved through "go routine")
package main

import (
	"fmt"
	"sync" //used for go routine
)

var waitGrp sync.WaitGroup

func main() {
//generally, main thread gets executed before the threads one and two. So output is not seen. To control it, use "wait" & "defer".
	
waitGrp.Add(2) //number of threads to be added into this group. Here indirectly "3" go routines are added ("main" is also a go routine).s


// waitGrp.Wait() // DEADLOCK, as the thread is in continuous waiting state (no decrements of wait group counter)

	for k:=1; k<=10; k++ {
		fmt.Println("main : ",k)
	}
	
	go func() { //"go" keyword is used to make that func as go routine
	
	//"defer" keyword: this func is executed in the end no matter what(use instead of Done()).
	//	defer waitGrp.Done() 

		for i:=101; i<=110; i++ {
		fmt.Println("one : ",i)
		}
		waitGrp.Done() //informs the wait group that this go routine is completed and wait group counter value is decremented by '1'
	}()

	go func() {
		//defer waitGrp.Done()

		for j:=11; j<=20; j++ {
		fmt.Println("two : ",j)
		}
		waitGrp.Done()
	}()

	waitGrp.Wait() //main thread will wait for the value to be zero.
}

//go run -race programname.go //displays race condition or deadlocks details