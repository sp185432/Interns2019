package main

import "fmt"

func main() {
var freesize_c, freesize_d, freesize_e int64
freesize_c = wmic logicaldisk C: get freespace
freesize_d = wmic logicaldisk D: get freespace
freesize_e = wmic logicaldisk E: get freespace
if(freesize_c < freesize_d) {
    if(freesize_c < freesize_e) {
        fmt.Printf("C drive has less space when compared to D and E")
    }else{
        if(freesize_d < freesize_e) {
        fmt.Printf("E drive has more space when compared to C and D")
        }else {
        fmt.Printf("D drive has more space when compared to C and E")
        }
    }

}

}