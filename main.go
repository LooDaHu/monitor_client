package main

import (
	"client/message"
	"client/utils"
	"time"
)

func main() {
	err := utils.InitConfig()
	if err != nil {
		return
	}
	utils.InitHTTPClient()
	utils.InitLogger()
	sysInfo := new(message.SystemInfo)
	for {
		go func() {
			info, err := sysInfo.GetSystemInfo()
			if err != nil {
				return
			}
			err = utils.HttpClient(info)
			if err != nil {
				return
			}
		}()
		time.Sleep(time.Minute)
	}
}
