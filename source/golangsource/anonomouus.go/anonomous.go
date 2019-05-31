package main
import(

	"fmt"
	"sync"
	//"runtime"
)
var waitGrp	sync.WaitGroup
func main() {
	waitGrp.Add(3)
	waitGrp.Wait()	
go func(){
defer waitGrp.Done()
for i := 0; i < 10; i++{
	fmt.Println("1 one function :",i)
}	
//waitGrp.Done()
}()
go func(){
	defer waitGrp.Done()
	for j := 0; j < 10; j++{
		fmt.Println("second function :",j)
	}	
//	waitGrp.Done()
}()
go func(){
	defer waitGrp.Done()
	for k := 0; k < 10; k++{
		fmt.Println("third function :",k)
	}	
	//waitGrp.Done()
	}()

}
