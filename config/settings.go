package config

import (
	"log"
	"os"

	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/utils"
	"gopkg.in/ini.v1"
)

type Server struct {
	AppMode  string
	HttpPort string
}

var ServerSetting = &Server{}

type Database struct {
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
}

var DatabaseSetting = &Database{}

type Redis struct {
	RdHost     string
	RdPort     string
	RdPassword string
}

var RedisSetting = &Redis{}

type Admin struct {
	Username string
	Password string
}

var AdminSetting = &Admin{}

var cfg *ini.File

// init 初始化读取配置文件
func init() {
	var err error
	cfg, err = ini.Load("./config/config_test.ini")

	if err != nil {
		log.Println("Open file error", err)
		os.Exit(1)
	}

	utils.MapTo(cfg, "server", ServerSetting)
	utils.MapTo(cfg, "database", DatabaseSetting)
	utils.MapTo(cfg, "redis", RedisSetting)
	utils.MapTo(cfg, "admin", AdminSetting)
}
