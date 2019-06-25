package main

import "fmt"

func main() {

	var a int
	var b int
	var c int
	fmt.Println("enter 2 numbers")
	fmt.Scanln(&a, &b)
	c = a + b
	fmt.Printf("Line 1 - Value of c is %d\n", c)

	c = a - b
	fmt.Printf("Line 2 - Value of c is %d\n", c)

	c = a * b
	fmt.Printf("Line 3 - Value of c is %d\n", c)

	c = a / b
	fmt.Printf("Line 4 - Value of c is %d\n", c)

	c = a % b
	fmt.Printf("Line 5 - Value of c is %d\n", c)

	a++
	fmt.Printf("Line 6 - Value of a is %d\n", a)

	a--
	fmt.Printf("Line 7 - Value of a is %d\n", a)
}
