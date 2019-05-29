//To know if a windows service is 'running' or not
package main

import (
	"fmt"
	"github.com/shirou/gopsutil/process"
	 )

func main() {
	var val *process.Process
	val.Pid = 2824
	ret, _ := val.IsRunning()
	fmt.Println(val.Status, ret)


/*	if() {
	fmt.Println("Service is Running")
	fmt.Println(pid, status)
	} else {
		fmt.Println("Service is NOT Running")
	}
*/
}