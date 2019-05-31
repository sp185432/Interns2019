package main

import "fmt"

func main() {
	//q := make(chan int) //unbuffered
	q := make(chan int,3)

	q <- 10
	q <- 20
	q <-30
	q<-40 //if it is included error arises deadlock as it exceeds the size
	//close(q)
	
	fmt.Println(<- q)
	fmt.Println(<- q)
	fmt.Println(<- q)
}
