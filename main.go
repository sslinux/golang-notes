package main

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	// "github.com/shirou/gopsutil/mem"  // to use v2
)

func main() {
	v, _ := mem.VirtualMemory()

	// almost every return value is a struct
	fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)

	fmt.Println()
	// convert to JSON. String() is also implemented
	fmt.Println(v)

	fmt.Println()
	c, _ := cpu.Info()
	fmt.Println(c)

	fmt.Println()

	parts, _ := disk.Partitions(true)
	fmt.Println(parts)

}
