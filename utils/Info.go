package utils

import (
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"

	. "net"
	"time"

	"client/message"
)

var hostIP IP

func InitHostIP() {
	hostIP = GetOutboundIP()
}

func GetSysInfo() message.SystemInfo {
	return message.SystemInfo{
		Host:    GetHostInfo(),
		Network: GetNetworkInfo(),
		CPU:     GetCPUInfo(),
		Memory:  GetMemInfo(),
		Disk:    GetDiskInfo(),
	}

}

func GetHostInfo() message.HostInfo {
	hostInfo, _ := host.Info()
	return message.HostInfo{
		HostName: hostInfo.Hostname,
		OS:       hostInfo.OS,
	}
}

func GetNetworkInfo() message.NetworkInfo {
	networkCounter, _ := net.IOCounters(false)

	return message.NetworkInfo{
		IPAddress:   hostIP.String(),
		BytesSend:   networkCounter[0].BytesSent,
		BytesRecv:   networkCounter[0].BytesRecv,
		PacketsSent: networkCounter[0].PacketsSent,
		PacketsRecv: networkCounter[0].PacketsRecv,
	}
}

func GetCPUInfo() message.CPUInfo {
	cpuInfo, _ := cpu.Info()
	logical, _ := cpu.Counts(true)
	physical, _ := cpu.Counts(true)
	percent, _ := cpu.Percent(time.Minute, false)

	return message.CPUInfo{
		ModelName: cpuInfo[0].ModelName,
		Logical:   logical,
		Physical:  physical,
		Percent:   percent[0],
	}
}

func GetMemInfo() message.MemoryInfo {
	memory, _ := mem.VirtualMemory()
	swap, _ := mem.SwapMemory()

	return message.MemoryInfo{
		Total: memory.Total,
		Used:  memory.Used,
		Swap:  swap.Total,
	}
}

func GetDiskInfo() message.DiskInfo {
	diskInfo, _ := disk.Usage("/")
	return message.DiskInfo{
		Total: diskInfo.Total,
		Used:  diskInfo.Used,
	}
}

func GetOutboundIP() IP {
	conn, err := Dial("udp", ServerAddress)
	if err != nil {
		SugarLogger.Error("unable to access server", err)
		panic("unable to access server, quit")
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*UDPAddr)

	return localAddr.IP
}
