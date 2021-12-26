package utils

import (
	"github.com/Unknwon/goconfig"
)

type GlobalCfg struct {
	AppId    string
	AppCode  string
	HostName string
}

var GlobalConfig GlobalCfg

func InitConfig() error {
	GlobalCfg, err := goconfig.LoadConfigFile("../conf/app.ini")
	if err != nil {
		SugarLogger.Error("Read Config Error @InitConfig", err)
		panic(ConfigReadError)
	}
	GlobalConfig.AppId, err = GlobalCfg.GetValue("app", "appid")
	GlobalConfig.AppCode, err = GlobalCfg.GetValue("app", "appcode")
	GlobalConfig.HostName, err = GlobalCfg.GetValue("host", "address")
	if err != nil {
		SugarLogger.Error("Get Config Value Error @InitConfig", err)
		return err
	}
	return nil
}
