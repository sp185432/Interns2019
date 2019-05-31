package main
import "fmt"

func main(){
	c:=make(chan int ,10)
	for i:=0; i<5 ;i++ { 
		c<-i
	}
	fmt.Println(len(c))
}