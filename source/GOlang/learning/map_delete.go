package main

import "fmt"

func main() {   
   /* create a map*/
   countryCapitalMap := map[string] string {"France":"Paris","Italy":"Rome","Japan":"Tokyo","India":"New Delhi"}
   
   fmt.Println("Original map")   
   
   /* print map */
   for country := range countryCapitalMap {
      fmt.Println("Capital of",country,"is",countryCapitalMap[country])
   }
   
   /* delete an entry */
   capital, ok := countryCapitalMap["France"]
   if (ok) {
	fmt.Println("Entry for france is found") 
	fmt.Println("Capital of France is", capital)
	delete(countryCapitalMap,"France")
	fmt.Println("Entry for France is deleted")  
	fmt.Println("Updated map")   
   }else{
	fmt.Println("Entry for France is not found")
   }
   
   
   /* print map */
   for country := range countryCapitalMap {
      fmt.Println("Capital of",country,"is",countryCapitalMap[country])
   }
}