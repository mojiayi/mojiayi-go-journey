package setting

import (
	"fmt"

	"gopkg.in/ini.v1"
)

var cfg *ini.File
var WebSetting = &WebConfig{}

func Setup() {
	var err error
	cfg, err = ini.Load("setting/my.ini")
	if err != nil {
		fmt.Println("failed while load setting file setting/my.ini,err: ", err)
	}

	mapToConfig("web", WebSetting)
}

func mapToConfig(section string, value interface{}) {
	err := cfg.Section(section).MapTo(value)
	if err != nil {
		fmt.Println("failed while cfg.MapTo "+section+",err: ", err)
	}
}

type WebConfig struct {
	Port        int
	ContextPath string
}
