package message

import (
	"client/utils"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
	. "net"
	"time"
)

type SystemInfo struct {
	Host    *HostInfo    `json:"host"`
	Network *NetworkInfo `json:"network"`
	CPU     *CPUInfo     `json:"cpu"`
	Memory  *MemoryInfo  `json:"memory"`
	Disk    *DiskInfo    `json:"disk"`
}

type HostInfo struct {
	HostName string `json:"host_name"`
	OS       string `json:"os"`
}

type NetworkInfo struct {
	IPAddress   string `json:"ip_address"`
	BytesSend   uint64 `json:"bytes_send"`
	BytesRecv   uint64 `json:"bytes_recv"`
	PacketsSent uint64 `json:"packets_sent"`
	PacketsRecv uint64 `json:"packets_recv"`
}

type CPUInfo struct {
	ModelName string  `json:"model_name"`
	Logical   int     `json:"logical"`
	Physical  int     `json:"physical"`
	Percent   float64 `json:"percent"`
}

type MemoryInfo struct {
	Total uint64 `json:"total"`
	Used  uint64 `json:"used"`
	Swap  uint64 `json:"swap"`
}

type DiskInfo struct {
	Total uint64 `json:"total"`
	Used  uint64 `json:"used"`
}

type Info interface {
	GetInfo() (interface{}, error)
}

type Summary interface {
	GetSystemInfo() (*SystemInfo, error)
}

func (s SystemInfo) GetSystemInfo() (*SystemInfo, error) {
	hostInfo, err := s.Host.GetInfo()
	if err != nil {
		return nil, err
	}
	networkInfo, err := s.Network.GetInfo()
	if err != nil {
		return nil, err
	}
	cpuInfo, err := s.CPU.GetInfo()
	if err != nil {
		return nil, err
	}
	memInfo, err := s.Memory.GetInfo()
	if err != nil {
		return nil, err
	}
	diskInfo, err := s.Disk.GetInfo()
	if err != nil {
		return nil, err
	}
	return &SystemInfo{
		Host:    hostInfo.(*HostInfo),
		Network: networkInfo.(*NetworkInfo),
		CPU:     cpuInfo.(*CPUInfo),
		Memory:  memInfo.(*MemoryInfo),
		Disk:    diskInfo.(*DiskInfo),
	}, nil
}

func (h HostInfo) GetInfo() (interface{}, error) {
	hostInfo, err := host.Info()
	if err != nil {
		utils.SugarLogger.Error("get host info error", err)
		return nil, err
	}
	return &HostInfo{
		HostName: hostInfo.Hostname,
		OS:       hostInfo.OS,
	}, nil
}

func (n NetworkInfo) GetInfo() (interface{}, error) {
	networkCounter, err := net.IOCounters(false)
	if err != nil || len(networkCounter) == 0 {
		utils.SugarLogger.Error("get network info error", err)
		return nil, err
	}
	hostIP := GetOutboundIP()
	return &NetworkInfo{
		IPAddress:   hostIP.String(),
		BytesSend:   networkCounter[0].BytesSent,
		BytesRecv:   networkCounter[0].BytesRecv,
		PacketsSent: networkCounter[0].PacketsSent,
		PacketsRecv: networkCounter[0].PacketsRecv,
	}, nil
}

func (c CPUInfo) GetInfo() (interface{}, error) {
	cpuInfo, err := cpu.Info()
	if err != nil || len(cpuInfo) == 0 {
		utils.SugarLogger.Error("get cpu info error", err)
		return nil, err
	}
	logical, err := cpu.Counts(true)
	if err != nil {
		utils.SugarLogger.Error("get cpu info error", err)
		return nil, err
	}
	physical, err := cpu.Counts(false)
	if err != nil {
		utils.SugarLogger.Error("get cpu info error", err)
		return nil, err
	}
	percent, err := cpu.Percent(time.Minute, false)
	if err != nil || len(percent) == 0 {
		utils.SugarLogger.Error("get cpu info error", err)
		return nil, err
	}

	return &CPUInfo{
		ModelName: cpuInfo[0].ModelName,
		Logical:   logical,
		Physical:  physical,
		Percent:   percent[0],
	}, nil
}

func (c MemoryInfo) GetInfo() (interface{}, error) {
	memory, err := mem.VirtualMemory()
	if err != nil {
		utils.SugarLogger.Error("get mem info error", err)
		return nil, err
	}
	swap, err := mem.SwapMemory()
	if err != nil {
		utils.SugarLogger.Error("get mem info error", err)
		return nil, err
	}
	return &MemoryInfo{
		Total: memory.Total,
		Used:  memory.Used,
		Swap:  swap.Total,
	}, nil
}

func (c DiskInfo) GetInfo() (interface{}, error) {
	diskInfo, err := disk.Usage("/")
	if err != nil {
		utils.SugarLogger.Error("get disk info error", err)
		return nil, err
	}
	return &DiskInfo{
		Total: diskInfo.Total,
		Used:  diskInfo.Used,
	}, nil
}

func GetOutboundIP() *IP {
	conn, err := Dial("udp", utils.GlobalConfig.HostName)
	if err != nil {
		utils.SugarLogger.Error("unable to access server", err)
		panic("unable to access server, quit")
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*UDPAddr)

	return &localAddr.IP
}
