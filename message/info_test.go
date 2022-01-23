package message

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCPUInfo_GetInfo(t *testing.T) {
	convey.Convey("TestCPUInfo_GetInfo_成功", t, func() {
		n := CPUInfo{}
		info, err := n.GetInfo()
		convey.So(info, convey.ShouldNotHaveSameTypeAs, CPUInfo{})
		convey.So(err, convey.ShouldEqual, nil)
	})
}

func TestDiskInfo_GetInfo(t *testing.T) {
	convey.Convey("TestDiskInfo_GetInfo_成功", t, func() {
		n := DiskInfo{}
		info, err := n.GetInfo()
		convey.So(info, convey.ShouldNotHaveSameTypeAs, DiskInfo{})
		convey.So(err, convey.ShouldEqual, nil)
	})
}

func TestGetOutboundIP(t *testing.T) {
	convey.Convey("TestGetOutboundIP_成功", t, func() {
		ip := GetOutboundIP()
		convey.So(ip, convey.ShouldNotBeNil)
	})
}

func TestHostInfo_GetInfo(t *testing.T) {
	convey.Convey("TestHostInfo_GetInfo_成功", t, func() {
		n := HostInfo{}
		info, err := n.GetInfo()
		convey.So(info, convey.ShouldNotHaveSameTypeAs, HostInfo{})
		convey.So(err, convey.ShouldEqual, nil)
	})
}

func TestMemoryInfo_GetInfo(t *testing.T) {
	convey.Convey("TestMemoryInfo_GetInfo_成功", t, func() {
		n := MemoryInfo{}
		info, err := n.GetInfo()
		convey.So(info, convey.ShouldNotHaveSameTypeAs, MemoryInfo{})
		convey.So(err, convey.ShouldEqual, nil)
	})
}

func TestNetworkInfo_GetInfo(t *testing.T) {
	convey.Convey("TestNetworkInfo_GetInfo_成功", t, func() {
		n := NetworkInfo{}
		info, err := n.GetInfo()
		convey.So(info, convey.ShouldNotHaveSameTypeAs, NetworkInfo{})
		convey.So(err, convey.ShouldEqual, nil)
	})
}

func TestSystemInfo_GetSystemInfo(t *testing.T) {
	convey.Convey("TestSystemInfo_GetSystemInfo_成功", t, func() {
		n := SystemInfo{}
		info, err := n.GetSystemInfo()
		convey.So(info, convey.ShouldNotHaveSameTypeAs, SystemInfo{})
		convey.So(err, convey.ShouldEqual, nil)
	})
}
