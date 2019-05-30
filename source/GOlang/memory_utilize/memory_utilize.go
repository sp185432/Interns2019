package main

import (
	"fmt"
	"strconv"

	"github.com/shirou/gopsutil/disk"
)

func main() {

	//vmStat, _ := mem.VirtualMemory()
	//fmt.Println("vmstat: %v ", vmStat)
	fmt.Printf("--------C DISK USAGE----------\n")
	diskStat, _ := disk.Usage("C:")

	fmt.Println("C drive total space in GB is: ", (diskStat.Total / 1024 / 1024 / 1024))
	fmt.Println("C drive Used space in GB is: ", (diskStat.Used / 1024 / 1024 / 1024))
	fmt.Println("C drive Free space in GB is: ", (diskStat.Free / 1024 / 1024 / 1024))

	fmt.Println("In C: drive Total disk space: " + strconv.FormatUint(diskStat.Total, 10) + " bytes")
	fmt.Println("Used disk space: " + strconv.FormatUint(diskStat.Used, 10) + " bytes")
	fmt.Println("Free disk space: " + strconv.FormatUint(diskStat.Free, 10) + " bytes")
	fmt.Println("Percentage disk space usage: " + strconv.FormatFloat(diskStat.UsedPercent, 'f', 2, 64))

	fmt.Printf("--------D DISK USAGE----------\n")
	diskStat1, _ := disk.Usage("D:")

	fmt.Println("D drive total space in GB is: ", (diskStat1.Total / 1024 / 1024 / 1024))
	fmt.Println("D drive Used space in GB is: ", (diskStat1.Used / 1024 / 1024 / 1024))
	fmt.Println("D drive Free space in GB is: ", (diskStat1.Free / 1024 / 1024 / 1024))

	//fmt.Println(diskStat1)

	fmt.Println("In D: drive Total disk space: " + strconv.FormatUint(diskStat1.Total, 10) + " bytes")
	fmt.Println("Used disk space: " + strconv.FormatUint(diskStat1.Used, 10) + " bytes")
	fmt.Println("Free disk space: " + strconv.FormatUint(diskStat1.Free, 10) + " bytes")
	fmt.Println("Percentage disk space usage: " + strconv.FormatFloat(diskStat1.UsedPercent, 'f', 2, 64))

	fmt.Printf("--------E DISK USAGE----------\n")

	diskStat2, _ := disk.Usage("E:")

	fmt.Println("E drive total space in GB is: ", (diskStat2.Total / 1024 / 1024 / 1024))
	fmt.Println("E drive Used space in GB is: ", (diskStat2.Used / 1024 / 1024 / 1024))
	fmt.Println("E drive Free space in GB is: ", (diskStat2.Free / 1024 / 1024 / 1024))

	//fmt.Println(diskStat2)

	fmt.Println("In E: drive Total disk space: " + strconv.FormatUint(diskStat2.Total, 10) + " bytes")
	fmt.Println("Used disk space: " + strconv.FormatUint(diskStat2.Used, 10) + " bytes")
	fmt.Println("Free disk space: " + strconv.FormatUint(diskStat2.Free, 10) + " bytes")
	fmt.Println("Percentage disk space usage: " + strconv.FormatFloat(diskStat2.UsedPercent, 'f', 2, 64))

}
