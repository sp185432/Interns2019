package main

import (
	"fmt"
	"os"
	"log"
	"bufio" //reading from console
	"io/ioutil" //for reading files
)

func readName() string {
	var fileName string = "" //path of file
	fmt.Println("Enter file name: ")
	fmt.Scanln(&fileName)
	return fileName
}

func file_exist(my_file string) (string) {
	fileInfo, stat_err := os.Stat(my_file) 
	if nil != stat_err {
		if os.IsNotExist(stat_err) {
		log.Println("File does not exist: ",stat_err)
		}
	}
	log.Println("File Stat is: ",fileInfo)
	return "exist"
}


func update(my_file string) string {
//opening file in write_only mode
fptr, open_err := os.OpenFile(my_file,os.O_WRONLY,0666)
if nil != open_err {
	log.Fatal("File Open Error for write mode: ",open_err)
}
log.Println("File is: ",fptr)

//Writing content to file
reader := bufio.NewReader(os.Stdin)
fmt.Println("Enter data into file, press '#' to end: ")
my_data, _ := reader.ReadString('#')
my_position, _ := fptr.Seek(-1,2)
fmt.Println(my_position)
fptr.WriteAt([]byte(my_data), my_position)
fptr.Close()
return "update"
}

func read(my_file string) string {
//opening file in read_only mode
fptr, open_err := os.OpenFile(my_file,os.O_RDONLY,0666)
if nil != open_err {
	log.Fatal("File Open Error for read mode: ",open_err)
}
log.Println("File is: ",fptr.Name())


//Reading contents of file
my_content, read_err := ioutil.ReadAll(fptr)
if nil != read_err {
	log.Fatal("File Read Error: ",read_err)
}
fmt.Println("File content is:")
fmt.Println(string(my_content)) 
fptr.Close()
return "read"
}

func main() {

	my_file := readName()

	val := file_exist(my_file)
	if "exist" == val {
		fmt.Println("File exist check done")
	} else {
		log.Fatalln("File exist check fail")
	}


	val = update(my_file)
	if "update" == val {
		fmt.Println("Successfully updated file")
	} else {
		fmt.Println("updation unsuccessful")
	}

	val = read(my_file)
	if "read" == val {
		fmt.Println("Successfully read file contents")
	} else {
		fmt.Println("file read unsuccessful")
	}
	
}