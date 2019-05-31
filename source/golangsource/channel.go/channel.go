package main

import (
  "fmt"
  "time"
)

func ping(c chan int) {
	
  for i := 0; ; i++ {
    c <- i
  }
}

 func print(c chan int) {
  for {
	  
	msg:=<- c
	
    fmt.Println(msg)
    time.Sleep(time.Second * 1)
  }
  fmt.Println(len(c))
}

func main() {
	var c chan int = make(chan int,11)
  //var cc chan int = make(chan int,5)
  go ping(c)
  go print(c)

  var input string
  fmt.Scanln(&input)
}
