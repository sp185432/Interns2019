package main

import "fmt"

func main() {
	var my_int int = 10
	var my_float float32 = 23.456
	//var my_char char = 'a'
	var my_string string = "navya"
	var my_complex complex64 = 12+5i
	var my_bool bool = true
	fmt.Println("Int= ",my_int)
	fmt.Println("Float= ",my_float)
//	fmt.Println("Character= "+my_char)
	fmt.Println("String= ",my_string)
	fmt.Println("Complex= ",my_complex)
	fmt.Println("Boolean= ",my_bool)

	var1 := 10
	var2 := 23.45667889
	var3 := "hello"
	var4 := 'a'
	var5 := 343+90i
	var6 := false
	var7 := '5'
	fmt.Printf("Types are: \n%T %T %T %T %T %T %T", var1, var2, var3, var4, var5, var6, var7)

}