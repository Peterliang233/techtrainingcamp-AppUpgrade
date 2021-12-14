package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/service/v1/rule"

	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/config"
	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/database/mysql"
	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/database/redis"
	v1 "github.com/Peterliang233/techtrainingcamp-AppUpgrade/router/v1"
)

func main() {
	router := v1.InitRouter()
	mysql.InitMysql()
	redis.InitRedis()

	server := &http.Server{
		Addr:           config.ServerSetting.HttpPort,
		Handler:        router,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Listen: %s\n", err)
		}
	}()

	// 开启一个定时器，定时将数据库的数据再次缓存到redis里面,防止redis挂了
	go rule.CacheData()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("ShutDown Server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Panic("Server Shutdown, ", err)
	}

	log.Println("Server Exiting...")
}
