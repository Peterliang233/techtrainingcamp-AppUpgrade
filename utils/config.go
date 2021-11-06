package utils

import (
	"gopkg.in/ini.v1"
	"log"
)

// MapTo 将配置文件里面的数据转化为结构体
func MapTo(cfg *ini.File, s string, i interface{}) {
	err := cfg.Section(s).MapTo(i)

	if err != nil {
		log.Fatalf("%s cfg Load error", err)
	}
}
