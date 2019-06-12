package main

import (
	"fmt"
	"net/http"
	"log"
)

//output: http://google.com --- 200 OK
//outptu: http://localhost:8080/blah" --- 404 Not Found

func mystat(url string) (string) {
	resp, err := http.Get(url) 
	if nil != err {
		log.Println("Error code is: ",err)
		return "NA"
	} else {
	return resp.Status
	}
}

func main() {
	url:=""
	fmt.Println("Enter a url: ")
	fmt.Scanln(&url)
	val := mystat(url)
	if "NA" == val {
		log.Panicln("Error in processing the url")
	} else {
	fmt.Println("Status is: ",val)
	}	
}

