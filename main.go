package main

import (
	"client/utils"
	"time"
)

func main() {
	utils.InitHostIP()
	utils.InitHTTPClient()
	utils.InitLogger()
	timeTicker := time.NewTicker(time.Minute)
	go func() {
		for {
			select {
			case <-timeTicker.C:
				sysInfo := utils.GetSysInfo()
				err := utils.HttpClient(sysInfo)
				if err != nil {
					continue
				}
			}
		}
	}()
}
