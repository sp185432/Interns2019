package main

import "fmt"

func main() {
	i:= 1
	switch i {
	case 5: fmt.Println("five")
	case 1: fmt.Println("one")
	case 2: fmt.Println("two")
	default: fmt.Println("no option")
	case 3,4: fmt.Println("three, four")
	}


	i= 1
	switch i {
	case 5: fmt.Println("-five")
	fallthrough //executes all the cases after the success case till the end of switch block
	case 1: fmt.Println("-one")
	fallthrough
	case 2: fmt.Println("-two")
	fallthrough
	default: fmt.Println("-no option")
	fallthrough
	case 6: fmt.Println("-six") //statements after this are not executed as 'fallthrough' is not included for this case.
	case 7: fmt.Println("-seven")
	case 3,4: fmt.Println("-three, four") //last statement doesnot need 'fallthrough' as it is the end statement.
	}

	
	switch i='g'; i {
	case 'a': fmt.Println("a")
	case 'b','c','d': fmt.Println("b, c, d")
	case 'e': fmt.Println("e")
	case 'f': fmt.Println("f")
	default: fmt.Println("no option")
	}

	switch i=3; { //no need to pass the variable into switch
		case 1 == i: fmt.Println("one")
		case 2 == i: fmt.Println("two")
		case 3 == i, 4 == i: fmt.Println("three, four")
		case 5 == i: fmt.Println("five")
		default: fmt.Println("no option")
	}

	j :="four" //since it is of different data type 'string', "i" cannot be used here as it is of type 'int'
	switch j {
		case "one": fmt.Println("one")
		case "two": fmt.Println("two")
		case "three": fmt.Println("three")
		case "four", "five": fmt.Println("four, five")
		default: fmt.Println("no option")
		}
	

}