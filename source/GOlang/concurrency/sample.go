package main

import (
	"fmt"
	"sync" //package
)

var waitGrp sync.WaitGroup //waitgroup is kind of function

func main() {

	waitGrp.Add(2) //any 2 threads

	go func() {
		defer waitGrp.Done()
		array := [2]int{1, 3}
		for i := 0; i < 500; i++ {

			fmt.Println("Inside first anonymous function: ", i)
			if i == 3 {
				array[i] = 10
			}
		}
		//waitGrp.Done()
	}()

	go func() {
		defer waitGrp.Done()
		for j := 501; j < 1000; j++ {
			fmt.Println("Inside second anonymous function: ", j)
		}
		//waitGrp.Done()
	}()

	waitGrp.Wait()

	fmt.Println("Inside main function")
}
