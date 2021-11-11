package main

import (
	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/config"
	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/database/mysql"
	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/database/redis"
	v1 "github.com/Peterliang233/techtrainingcamp-AppUpgrade/router/v1"
)

func main() {
	router := v1.InitRouter()

	mysql.InitMysql()
	redis.InitRedis()

	if err := router.Run(config.ServerSetting.HttpPort); err != nil {
		panic(err)
	}
}
