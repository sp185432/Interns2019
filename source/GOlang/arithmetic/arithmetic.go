package main

import "fmt"

func add(num1 int, num2 int) (ret int) { //using "ret" keyword
	ret = num1 + num2
	return
}

func subtract(num1 int, num2 int) (int) {
	return num1 - num2
}

func multiply(num1 int, num2 int) (int) {
	return num1 * num2
}

func divide(num1 int, num2 int) (int) {
	if(0 == num2) {
		fmt.Println("divide by zero exception occured")
		return -1
	}
	return num1 / num2
}

func main() {
	num1, num2 := 0, 0
	fmt.Print("enter two numbers: ")
	fmt.Scanln(&num1, &num2)

	for{
		operation := 0
		fmt.Println("enter choice: 1.add, 2.subtract, 3.multiply, 4.divide, 5.EXIT")
		fmt.Scanln(&operation)

		switch operation {
			case 1:fmt.Println("sum= ",add(num1, num2))
			case 2:fmt.Println("subtraction= ",subtract(num1, num2))
			case 3:fmt.Println("multiplication= ",multiply(num1, num2))
			case 4: value := divide(num1, num2)
			if(-1 != value) {
				fmt.Println("quotent= ",value)
			}
			case 5:break
			default:fmt.Println("invalid choice, choose again")
		}//switch
		if(5 == operation) {
			break
		}//if
	}//for

	fmt.Println("program ended")
}