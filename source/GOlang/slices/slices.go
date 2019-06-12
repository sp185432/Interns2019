package main

import "fmt"

func main() {
	
	//identifying array and slice: array is declared with size([5]) and slice is declared with empty size([]).
	sl := []int{} //empty slice
	fmt.Println(sl)
	sl1 := make([]int, 5) //length=5
	fmt.Println(sl1)
	sl2 := make([]int, 3, 5) //length=3, capacity=5. Condition: length<=capacity
	sl2[0] = 11
	sl2[2] = 33
	//sl2[4] = 55 //not possible as len=3, only appending is possible
	sl2 = append(sl2,44)
	sl2 = append(sl2,55,66)
	fmt.Println(sl2, "Length = ",len(sl2), " Capacity = ",cap(sl2))
	
	sl3 := []int{11,22,33}
	fmt.Println(sl3, "Length = ",len(sl3))
	sl4 := []int{1:102, 4: 105}
	fmt.Println(sl4)
	sl3 = sl4 //possible to assign slice of different (<, >, =) length to another slice
	fmt.Println(sl3)
	fmt.Println(sl4)
	sl3 = sl4[1:4] //[0:n] slices till n-1
	fmt.Println(sl3, "Length = ",len(sl3)," Capacity = ",cap(sl3))
	sl5 := []string{"blah", "bla"}
	fmt.Println(sl5)
	sl6 := []int{1:2, 3:4}
	fmt.Println(sl6)
	sl6 = append(sl6,6)
	fmt.Println(sl6)
	
}