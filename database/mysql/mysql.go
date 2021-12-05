package mysql

import (
	"fmt"
	"log"
	"time"

	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/model"

	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	Db *gorm.DB
)

// InitMysql 初始化数据库
func InitMysql() {
	var err error

	Db, err = gorm.Open(config.DatabaseSetting.Db,
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			config.DatabaseSetting.DbUser,
			config.DatabaseSetting.DbPassword,
			config.DatabaseSetting.DbHost,
			config.DatabaseSetting.DbPort,
			config.DatabaseSetting.DbName,
		))

	if err != nil {
		log.Println("数据库打开错误", err)
	}

	Db.SingularTable(true)

	Db.AutoMigrate(model.Device{})
	Db.AutoMigrate(model.Rule{})
	Db.AutoMigrate(model.User{})
	Db.AutoMigrate(model.RuleState{})

	Db.DB().SetMaxIdleConns(10)

	Db.DB().SetMaxOpenConns(100)

	Db.DB().SetConnMaxLifetime(10 * time.Second)
}
