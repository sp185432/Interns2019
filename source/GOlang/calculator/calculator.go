package arithmetic

import "fmt"

//Addition function to perform addition of entered two numbers
func Addition(num1, num2 int) (result int) {
	result = num1 + num2
	//fmt.Printf("Addition_result: %d num1: %d num2: %d\n", result, num1, num2)
	return result
}

//Subtraction function to perform subtraction of entered two numbers
func Subtraction(num1, num2 int) (result int) {
	result = num1 - num2
	//fmt.Printf("Subtraction_result: %d num1: %d num2: %d\n", result, num1, num2)
	return result
}

//Multiplication function to perform multiplication of entered two numbers
func Multiplication(num1, num2 int) (result int) {
	result = num1 * num2
	//fmt.Printf("Multiplication_result: %d num1: %d num2: %d\n", result, num1, num2)
	return result
}

//Division function to perform division of entered two numbers
func Division(num1, num2 int) (result int) {
	if num2 == 0 {
		fmt.Printf("Divide by zero is not possible. Denominator should be a non zero value\n")
	} else {
		result = num1 / num2
		//fmt.Printf("Division_result: %d num1: %d num2: %d\n", result, num1, num2)
	}

	return result

}

/*
func main() {
	var num1, num2, choice int
	fmt.Printf("Enter number 1: ")
	fmt.Scan(&num1)
	fmt.Printf("Enter number 2: ")
	fmt.Scan(&num2)
	fmt.Printf("\nChoose the operation that you want to perform.\n")
	fmt.Printf("1.Addition\n 2. Subtraction\n 3. Multiplication\n 4. Division\n 5. Exit\n")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		Addition(num1, num2) //calling the function Addition

	case 2:
		Subtraction(num1, num2) //calling the function Subtraction

	case 3:
		Multiplication(num1, num2) //calling the function Multiplication

	case 4:
		Division(num1, num2) //calling the function Division
	case 5:
		break
	default:
		fmt.Printf("\nEnter a valid number")
	}
}
*/
