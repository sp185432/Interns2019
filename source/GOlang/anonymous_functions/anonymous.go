package main

import "fmt"

func main() {
	func() {
		fmt.Println("Hello")
	}()

	func() {
		fmt.Println("World!")
	}()
	func() {
		fmt.Println("WELCOME!")
	}()
}
