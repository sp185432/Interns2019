package main

import "fmt"

func main() {
	
	//empty map
	map1 := make(map[int]int) //map[key]value
	fmt.Println(map1)

	map2 := map[int]int{1:11, 2:22, 3:33}
	fmt.Println(map2)

	map3 := map[string]int{"abc":1, "def":2, "ghi":3}
	fmt.Println(map3, "Length = ",len(map3))
	fmt.Println(map2[2])
	delete(map2,2)
	fmt.Println(map2)

}