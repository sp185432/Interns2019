package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	// robocopy //BLAH/readfile/datatypes.exe C:/Users/Navya
	cmd := exec.Command("robocopy", "//BLAH/readfile", "C:/Users/Navya")
	err := cmd.Run()
	if nil != err {
		log.Printf("Error code is: %T\n", err)
	}

	fmt.Println(cmd.Output())
	cmd = exec.Command("mv", "C:/Users/Navya/datatypes.exe", "C:/Users/Navya/go/src/exeserver")
	err = cmd.Run()
	if nil != err {
		log.Printf("Error code is: %T\n", err)
	}

	fmt.Println(cmd.Output())

	fmt.Println(cmd.Output())
	out, outputErr := exec.Command("./datatypes.exe").Output()
	//err = cmd.Run()
	if nil != outputErr {
		log.Printf("Error code is: %T\n", outputErr)
	}

	fmt.Println(string(out))

}
