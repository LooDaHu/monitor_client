package main

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
	// "github.com/shirou/gopsutil/mem"  // to use v2
)

func main() {
	//v, _ := mem.VirtualMemory()
	//sysInfo, _ := host.Info()
	//cpuInfo, _ := cpu.Percent(time.Second, true)
	cpuInfo, _ := cpu.Counts(true)
	cpuInfo1, _ := cpu.Counts(false)
	network, _ := net.IOCounters(false)
	memory, _ := mem.VirtualMemory()
	swap, _ := mem.SwapMemory()
	// almost every return value is a struct
	//fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)
	//fmt.Printf("", sysInfo)
	fmt.Printf("%v \n", cpuInfo)
	fmt.Printf("%v \n              ", cpuInfo1)
	fmt.Printf("%v \n              ", network)
	fmt.Printf("%v \n              ", memory)
	fmt.Printf("%v \n              ", swap)
	// convert to JSON. String() is also implemented
	//fmt.Println(v)
	// https://blog.csdn.net/weixin_30485799/article/details/98980008
}
