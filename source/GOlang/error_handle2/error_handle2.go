package main

import (
	"fmt"
	"sync"
	"os"
	//"log" //for fatal, panic etc. messages display
)

var waitGrp sync.WaitGroup

func myfun(temp int) {
	
	defer waitGrp.Done()

	fmt.Println("My Function: ",temp)
	fptr, err := os.Open("file.txt")
	if nil != err {
		fmt.Println("Error is: ",err)
	//	log.Fatalln("Error is: ",err) //used for login purpose with timestamp
	//	log.Panic("Error is: ",err)	
	
		return	
	}
	
	fmt.Println(fptr.Name()," open is successful in thread ",temp)
	fptr.Close()
	
	//waitGrp.Done() //causes deadlock in this code because when error is caused, the control returns to main in the previous step and does not execude Done(). The main thread keeps on waiting for the thread counter to be decremented.
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