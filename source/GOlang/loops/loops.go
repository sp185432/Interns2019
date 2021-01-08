package main

import "fmt"

func main() {

	//for loop
	for i:=0; i<10; i++ {
		if(6 == i) {
			break
		}
		fmt.Println(i)
	}

	goto my_flag
	
/*	fmt.Println("hello") //unreachable code warning because of skipping due to go-to statement
	fmt.Println("blah") //unreachable code warning because of skipping due to go-to statement
*/
	my_flag: fmt.Println("go-to successful")

	for i:=100; i<=105; i++ {
		if(13 == i) {
			continue
		}
		fmt.Println(i)
	}

	//while
	i:= 25
	for i>20 {
		fmt.Println(i)
		i--
	}

	arr := [5]int{10,20,30,40,50}
	for i,v := range arr {
		fmt.Println(i,v)
	}

}