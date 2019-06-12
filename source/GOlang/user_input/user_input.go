package main

import "fmt"

func main() {
	var one,two int = 0, 0
	fmt.Println("Enter 2 values: ")
	fmt.Scanln(&one, &two) //values must be entered in a single line
	fmt.Println("The values are: ",one,two)
	fmt.Println("Enter 2 values: ")
	fmt.Scan(&one, &two) //values can be entered in different lines also
	fmt.Println("Sum = ",one+two)
}