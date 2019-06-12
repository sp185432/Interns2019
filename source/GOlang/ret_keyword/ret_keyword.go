package main

import "fmt"

func myfun(n1 int, n2 int) (ret int) {
    ret = n1 + n2
    return
}

func main() {
    v1:= myfun(3,4)
    fmt.Println(v1)
}