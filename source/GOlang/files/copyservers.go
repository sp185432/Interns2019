package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("robocopy", "//BLAH/readfile", "C:/Users/Navya")
	err := cmd.Run()
	if nil != err {
		log.Printf("Error code is: %T\n", err)
	}

	fmt.Println(cmd.Output())

}
