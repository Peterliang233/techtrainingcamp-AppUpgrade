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
	DbPassword string
	DbName     string
}

var DatabaseSetting = &Database{}

type Redis struct {
	RdHost     string
	RdPort     string
	RbPassword string
}

var RedisSetting = &Redis{}

var cfg *ini.File

// Init 初始化读取配置文件
func Init() {
	var err error
	cfg, err = ini.Load("./config/config.ini")

	if err != nil {
		log.Println("Open file error", err)
		os.Exit(1)
	}

	utils.MapTo(cfg, "server", ServerSetting)
	utils.MapTo(cfg, "database", DatabaseSetting)
	utils.MapTo(cfg, "redis", RedisSetting)
}