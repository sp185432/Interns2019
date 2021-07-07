package main

import "fmt"

func main() {
	messages := make(chan string, 2)

	messages <- "sharanya" //sending
	messages <- "gandla"
	//messages <- "gandla"

	fmt.Println(<-messages) //receiving
	fmt.Println(<-messages)
}
