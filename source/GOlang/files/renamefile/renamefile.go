package main

import (
	"fmt"
	"os"
	"log"
)

func my_rename(old_name string, new_name string) (string) {
	rename_err := os.Rename(old_name, new_name)
	if nil != rename_err {
		log.Fatal("File Open Error for write mode: ",rename_err)
	}
	
	//opening file in read_only mode
	fptr, open_err := os.OpenFile(new_name,os.O_RDONLY,0666)
	if nil != open_err {
		log.Fatal("File Open Error for read mode: ",open_err)
	}
	log.Println("File Name after renaming is: ",fptr.Name())
	fptr.Close()
	return "yes"
}

func file_exist(old_name string) (string) {
	fileInfo, stat_err := os.Stat(old_name) 
	if nil != stat_err {
		if os.IsNotExist(stat_err) {
		log.Println("File does not exist: ",stat_err)
		}
	}
	log.Println("File Stat is: ",fileInfo)
	return "exist"
}

func main() {
	var old_name, new_name string = "","" //paths of files
	fmt.Println("Enter name of file: ")
	fmt.Scanln(&old_name)

	val := file_exist(old_name)
	if "exist" == val {
		fmt.Println("File exist check done")
	} else {
		log.Fatalln("File exist check fail")
	}

	fmt.Println("Enter new name of file: ")
	fmt.Scanln(&new_name)
	if("" == new_name) {
		log.Fatal("invalid file name")
	}

	val= my_rename(old_name, new_name)
	if "yes" == val {
		fmt.Println("Successfully renamed file")
	} else {
		fmt.Println("renaming unsuccessful")
	}

}