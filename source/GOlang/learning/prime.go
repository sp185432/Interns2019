package main

import "fmt"

func main() {
	var number int
	var flag int
    fmt.Printf("Enter the number: ")
    fmt.Scanf("%d", &number)
    for i:=1; i<= number;i++ {
        if(number % i == 0) {
			flag++
		}
    }
    if(flag == 2) {
	fmt.Printf("\nThe entered number %d is a Prime number", number)
	}else {
	fmt.Printf("\nThe entered number %d is not a Prime number", number)
	}
}