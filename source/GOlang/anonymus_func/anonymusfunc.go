package main

import "fmt"

func main() {
	func() {
		fmt.Println("one")
	}()

	func() {
		fmt.Println("two")
	}()
}