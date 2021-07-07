package main

import "fmt"

var glob_vari int
var test_var int = 100

func main() {
   /* local variable declaration */
   var a, b, local_vari int 

   /* actual initialization */
   a = 10
   b = 20

   var test_var int = 200

   local_vari = a + b
   glob_vari = a + b
   fmt.Printf ("value of a = %d, b = %d and local_vari = %d\n", a, b, local_vari)
   fmt.Printf ("\n value of a = %d, b = %d and glob_vari = %d\n", a, b, glob_vari)
   fmt.Printf("\n value of test_var is %d",test_var)
}