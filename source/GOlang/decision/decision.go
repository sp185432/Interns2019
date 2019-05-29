package main
import "fmt"
func main() {

	var myvar = false

	if(myvar) {
	fmt.Println("true")
	} else {
	fmt.Println("false")
	}

	my_var := 10
	if(my_var < 10) {
		fmt.Println("less than 10")
	} else if(10 == my_var) {
		fmt.Println("equals to 10")
	}else {
		fmt.Println("greater than 10")
	}
}