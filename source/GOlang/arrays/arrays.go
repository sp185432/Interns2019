package main

import "fmt"

func main() {
	
	arr1 := [5]int{1,2,3,4,5}
	fmt.Println(arr1)
	var arr2 [5]int //sets elements to zero by default
	fmt.Println(arr2)
	arr2[0] = 11
	arr2[3] = 33
	fmt.Println(arr2)
	arr3 := [...]int{1,2,3,4,5,6,7} //size is set automatically by the compiler depending on the number of elements initialized. Initialization is compulsory.
	fmt.Println(arr3, "Length = ",len(arr3))
	arr4 := [5]int{1:102, 4: 105}
	fmt.Println(arr4)
	//arr4 = arr3 // not possible as sizes of both arrays are different. Both sizes MUST be EQUAL to assign.
	arr4 = arr1
	fmt.Println(arr4)
	arr5 := [2]string{"blah", "bla"}
	fmt.Println(arr5)
	arr6 := [2]int32{'1', '2'}
	fmt.Println(arr6)
	

	
}