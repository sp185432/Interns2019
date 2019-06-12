//To know if a windows service is 'running' or not
package main

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	 )

func main() {
	//cpu details
	cpuStat, err := cpu.Info()
	if nil != err {
		fmt.Println("Error: ",err)
	}
	fmt.Println("CPU Stat: ",cpuStat)
	//fmt.Printf("Type of cpustat is: %T\n",cpuStat)

	//cpu utilization percentage
	percentage, err := cpu.Percent(0, true)
	if nil != err {
		fmt.Println("Error: ",err)
	}
	fmt.Println("CPU %: ",percentage)
/*	if() {
	fmt.Println("Service is Running")
	fmt.Println(pid, status)
	} else {
		fmt.Println("Service is NOT Running")
	}
*/
}