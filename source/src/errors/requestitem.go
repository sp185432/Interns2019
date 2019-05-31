package main

import (
	"errors"
	"log"
)

var errNotFound = errors.New("Item not found")

func main() {
	err := getItem(123) // This would throw errNotFound
	if err != nil {
		switch err {
		case errNotFound:
			log.Println("Requested item not found")
		default:
			log.Println("Unknown error occurred")
		}
	}
}
