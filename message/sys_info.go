package message

type SystemInfo struct {
	Host    HostInfo    `json:"host"`
	Network NetworkInfo `json:"network"`
	CPU     CPUInfo     `json:"cpu"`
	Memory  MemoryInfo  `json:"memory"`
	Disk    DiskInfo    `json:"disk"`
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
