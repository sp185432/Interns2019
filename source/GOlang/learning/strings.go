package main

import (
	"fmt" 
	"strings")

func main() {
   var greeting =  "Hello world!"
   
   fmt.Printf("String Length is: ")
   fmt.Println(len(greeting))  
   greetings :=  []string{"Welcome","Interns!"}   
   fmt.Println(strings.Join(greetings, " "))
}