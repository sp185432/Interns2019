package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

var (
	fileInfo *os.FileInfo
	err      error
	newFile  *os.File
)

func main() {
	// Stat returns file info. It will return
	// an error if there is no file.
	var fileName string
	fmt.Println("enter file name")
	fmt.Scanln(&fileName)
	fileInfo, err := os.Open(fileName + ".txt")
	if err != nil {
		if os.IsNotExist(err) {
			newFile, err = os.Create(fileName + ".txt")
			newFile.WriteString("hello world")
			if err != nil {
				log.Fatal(err)
			}
			log.Println(newFile)
			newFile.Close()
		}
	} else {
		buf := new(bytes.Buffer)
		buf.ReadFrom(fileInfo)
		contents := buf.String()

		fmt.Print(contents)
		log.Println("File does exist. File information:")
		log.Println(fileInfo)
	}
}
