package main

import "fmt"

func main() {
	//untyped constants
	const untyped_int = 100 //not assigning any data type
	fmt.Println(untyped_int)

	var my_var int = 0
	my_var = untyped_int
	fmt.Println(my_var)

	//typed constants
	const typed_int int = 10
	fmt.Println(typed_int)

	var my_int int = 0 
	my_int = typed_int
	fmt.Println(my_int)

	var my_float float32 = 0
	//my_float = typed_int //error as variable of other data type cannot be assigned with other data typed constant 
	fmt.Println(my_float)

}
