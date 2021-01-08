package main

import (
	"bufio" //reading from console
	"fmt"
	"io/ioutil" //for reading files
	"log"
	"os"
)



func readName() string {
	var fileName string = "" //path of file
	fmt.Println("Enter file name: ")
	fmt.Scanln(&fileName)
	if "" == fileName {
		return "invalidname"
	}
	return fileName
}

func createFile(my_file string) string {
	//Creating a new file
	fptr, create_err := os.OpenFile(my_file, os.O_CREATE, 0666)
	if nil != create_err {
		log.Fatal("File Creation Error: ", create_err)
	}
	log.Println("File is: ", fptr)
	fptr.Close()
	return "create"
}

func writeFile(my_file string) string {
	//Opening file in write_only mode
	fptr, open_err := os.OpenFile(my_file, os.O_WRONLY, 0666)
	if nil != open_err {
		log.Fatal("File Open Error for write mode: ", open_err)
	}
	log.Println("File is: ", fptr)

	//Writing content to file
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter data into file, press '#' to end: ")
	my_data, _ := reader.ReadString('#')

	fptr.WriteString(my_data)
	fptr.Close()
	return "write"
}

func readFile(my_file string) string {
	//Opening file in read_only mode
	fptr, open_err := os.OpenFile(my_file, os.O_RDONLY, 0666)
	if nil != open_err {
		log.Fatal("File Open Error for read mode: ", open_err)
	}
	log.Println("File is: ", fptr.Name())

	//Reading contents of file
	my_content, read_err := ioutil.ReadAll(fptr)
	if nil != read_err {
		log.Fatal("File Read Error: ", read_err)
	}
	fmt.Println("File content is:")
	fmt.Println(string(my_content))
	fptr.Close()
	return "read"
}

func main() {
	
	my_file := readName()
	if "invalidname" != my_file {
		fileInfo, stat_err := os.Stat(my_file)
		if nil != stat_err {
			if os.IsNotExist(stat_err) { //execute the program only if the file doesnot exist

				val := createFile(my_file)
				if "create" == val {
					fmt.Println("Successfully created file")
				} else {
					fmt.Println("file creation unsuccessful")
				}

				val = writeFile(my_file)
				if "write" == val {
					fmt.Println("Successfully written file")
				} else {
					fmt.Println("file write unsuccessful")
				}

				val = readFile(my_file)
				if "read" == val {
					fmt.Println("Successfully read file contents")
				} else {
					fmt.Println("file read unsuccessful")
				}
			}
		} else { //file exists case

			log.Fatal("File already exists. File info is: ", fileInfo)
		}
	} else {
		log.Fatal("Invalid file name")
	}
}
