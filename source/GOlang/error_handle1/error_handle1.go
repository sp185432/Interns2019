package main

import (
	"fmt"
	"sync"
)

var waitGrp sync.WaitGroup

func myfun(temp int) {
	
	fmt.Println("My Function: ",temp)
	var arr [3]int
	
	//This code causes "index out of range" error at run-time and abrupt termination of program
	// for i:=0; i<4; i++ { 
	// 	arr[i] = i
	// 	fmt.Println(arr[i])
	// }

	for i:=0; i<4; i++ { 
	if(i >= 3) {  //handling error
		fmt.Println("Array index out of bound, terminating entry of values into array")
		break
	}
		arr[i] = i
	}
	fmt.Println("Array-",temp,": ",arr)
	waitGrp.Done()
	
}

func main() {

	fmt.Println("MAIN")
	waitGrp.Add(3)

	for i:=0; i<3; i++ {
		go myfun(i)
	}

	waitGrp.Wait()
	fmt.Println("end")
}